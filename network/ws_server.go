package network

import (
	"crypto/tls"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"github.com/cloudfreexiao/antnet/logger"
)

type WSServer struct {
	Addr string
	MaxConnNum int
	PendingWriteNum int
	MaxMsgLen       uint32
	HTTPTimeout     time.Duration
	CertFile        string
	KeyFile         string
	NewAgent        func(*WSConn) Agent
	ln              net.Listener
	handler         *WSHandler
}

type WSHandler struct {
	maxConnNum      int
	pendingWriteNum int
	maxMsgLen       uint32
	newAgent        func(*WSConn) Agent
	upgrader        websocket.Upgrader
	conns           WSConnSet
	mutexConns      sync.Mutex
	wg              sync.WaitGroup
}

func (handler *WSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ERROR(w, "Method not allowed", 405)
		return
	}
	conn, err := handler.upgrader.Upgrade(w, r, nil)
	if err != nil {
		DEBUG("upgrade error:", Inspect(err))
		return
	}
	conn.SetReadLimit(int64(handler.maxMsgLen))

	handler.wg.Add(1)
	defer handler.wg.Done()

	handler.mutexConns.Lock()
	if handler.conns == nil {
		handler.mutexConns.Unlock()
		conn.Close()
		return
	}
	if len(handler.conns) >= handler.maxConnNum {
		handler.mutexConns.Unlock()
		conn.Close()
		INFO("too many connections")
		return
	}
	handler.conns[conn] = struct{}{}
	handler.mutexConns.Unlock()

	wsConn := newWSConn(conn, handler.pendingWriteNum, handler.maxMsgLen)
	agent := handler.newAgent(wsConn)
	agent.Run()

	// cleanup
	wsConn.Close()
	handler.mutexConns.Lock()
	delete(handler.conns, conn)
	handler.mutexConns.Unlock()
	agent.OnClose()
}

func (server *WSServer) Start() {
	ln, err := net.Listen("tcp", server.Addr)
	if err != {
		ERROR("WS listen err", Inspect(err))
	}

	if server.MaxConnNum <= 0 {
		server.MaxConnNum = 100
		INFO("WS invalid MaxConnNum, reset to:", Inspect(server.MaxConnNum) )
	}

	if server.PendingWriteNum <= 0 {
		server.PendingWriteNum = 100
		INFO("WS invalid PendingWriteNum, reset to:", Inspect(server.PendingWriteNum))
	}

	if server.MaxMsgLen <= 0 {
		server.MaxMsgLen = 4096
		INFO("WS invalid MaxMsgLen, reset to:", Inspect(server.MaxMsgLen) )
	}

	if server.HTTPTimeout <= 0 {
		server.HTTPTimeout = 10 * time.Second
		INFO("WS invalid HTTPTimeout, reset to:", Inspect(server.HTTPTimeout) )
	}

	if server.NewAgent == nil {
		ERROR("NewAgent must not be nil")
	}

	if server.CertFile != "" || server.KeyFile != "" {
		config := &tls.Config{}
		config.NextProtos = []string{"http/1.1"}

		var err error
		config.Certificates = make([]tls.Certificate, 1)
		config.Certificates[0], err = tls.LoadX509KeyPair(server.CertFile, server.KeyFile)
		if err != nil {
			ERROR("WS Certificates error:", Inspect(err))
		}

		ln = tls.NewListener(ln, config)
	}

	server.ln = ln

	server.handler = &WSHandler{
		maxConnNum:      server.MaxConnNum,
		pendingWriteNum: server.PendingWriteNum,
		maxMsgLen:       server.MaxMsgLen,
		newAgent:        server.NewAgent,
		conns:           make(WSConnSet),
		upgrader: websocket.Upgrader{
			HandshakeTimeout: server.HTTPTimeout,
			CheckOrigin:      func(_ *http.Request) bool { return true },
		},
	}

	httpServer := &http.Server{
		Addr:           server.Addr,
		Handler:        server.handler,
		ReadTimeout:    server.HTTPTimeout,
		WriteTimeout:   server.HTTPTimeout,
		MaxHeaderBytes: 1024,
	}

	go httpServer.Serve(ln)
}

func (server *WSServer) Close() {
	server.ln.Close()

	server.handler.mutexConns.Lock()
	for conn := range server.handler.conns {
		conn.Close()
	}
	server.handler.conns = nil
	server.handler.mutexConns.Unlock()

	server.handler.wg.Wait()
}