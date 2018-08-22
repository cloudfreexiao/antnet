package network

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/cloudfreexiao/antnet/logger"
)

type TCPClient struct {
	sync.Mutex
	Addr string
	ConnNum int
	ConnectInterval time.Duration
	PengingWriteNum int
	AutoReconnect bool
	NewAgent func(*TCPConn) Agent
	conns ConnSet
	wg sync.WaitGroup
	closeFlag bool

	//msg parser
	LenMsgLen int
	MinMsgLen uint32
	MaxMsgLen uint32
	LittleEndian bool
	msgParser *MsgParser
}

func (client *TCPClient) Start()  {
	DEBUG("++++TCPClient Init++++", client.ConnNum)
	client.init()
	DEBUG("++++TCPClient Will Conn++++", client.ConnNum)

	for i := 0; i< client.ConnNum; i++ {
		client.wg.Add(1)
		go client.connect()
	}
}

func (client *TCPClient) init()  {
	client.Lock()
	defer client.Unlock()

	if client.ConnNum <= 0 {
		client.ConnNum = 1
		DEBUG("invalid ConnNum, reset to ", client.ConnNum )
	}

	if client.ConnectInterval <= 0 {
		client.ConnectInterval = 3 * time.Second
		DEBUG("invalid ConnectInterval, reset to ", client.ConnectInterval)
	}

	if client.PendingWriteNum <= 0 {
		client.PendingWriteNum = 100
		DEBUG("invalid PendingWriteNum, reset to ", client.PendingWriteNum)
	}

	if client.NewAgent == nil {
		ERROR("TCPClient NewAgent Must Not Be Nil")
	}

	if client.conns != nil {
		ERROR("TCPClient is Running")
	}

	client.conns = make(ConnSet)
	client.closeFlag = false

	// msg parser
	msgParser := NewMsgParser()
	msgParser.SetMsgLen(client.LenMsgLen, client.MinMsgLen, client.MaxMsgLen)
	msgParser.SetByteOrder(client.LittleEndian)
	client.msgParser = msgParser
}

func (client *TCPClient) dial() net.Conn {
	for {
		conn, err := net.Dial("tcp", client.Addr)
		if err == nil || client.closeFlag {
			return conn
		}

		INFO("connect to %v error: %v", client.Addr, err)
		time.Sleep(client.ConnectInterval)
		continue
	}
}

func (client *TCPClient) connect()  {
	defer client.wg.Done()
	DEBUG("++++++++++ tcp connect+++++++")

reconnect:
	conn := client.dial()
	if conn == nil {
		DEBUG("++++++++++ tcp dial failed+++++++")
		return 
	}

	client.Lock()
	if client.closeFlag {
		client.Unlock()
		conn.Close()
		DEBUG("++++++++++ tcp will close+++++++")
		return
	}

	client.conns[conn] = struct{}{}
	client.Unlock()
	DEBUG("++++++++++ tcp connected+++++++")

	tcpConn := newTCPConn(conn, client.PengingWriteNum, client.msgParser)
	agent := client.NewAgent(tcpConn)
	agent.Run()

	// clean up
	tcpConn.Close()
	client.Lock()
	delete(client.conns, conn)
	client.Unlock()
	agent.OnClose()

	if client.AutoReconnect {
		time.Sleep(client.ConnectInterval)
		goto reconnect
	}
}

