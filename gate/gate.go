package gate

import (
	_ "net"
	_ "reflect"
	"time"


	"github.com/cloudfreexiao/antnet/logger"
	"github.com/cloudfreexiao/antnet/network"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type IGateService interface {
	GetAgentActor(Agent) (*actor.PID, error)
}

type Gate struct {
	MaxConnNum int
	PendingWriteNum int
	MaxMsgLen       uint32
	Processor       network.Processor

	// ws
	WSAddr      string
	HTTPTimeout time.Duration
	CertFile    string
	KeyFile     string
	
	// tcp
	TCPAddr      string
	LenMsgLen    int
	LittleEndian bool

	wsServer  *network.WSServer
	tcpServer *network.TCPServer
}

func (gate *Gate) Run(gs IGateService) {
	var wsServer *network.WSServer
	if gate.WSAddr != "" {
		wsServer = new(network.WSServer)
		wsServer.Addr = gate.WSAddr
		wsServer.MaxConnNum = gate.MaxConnNum
		wsServer.PendingWriteNum = gate.PendingWriteNum
		wsServer.HTTPTimeout = gate.HTTPTimeout
		wsServer.CertFile = gate.CertFile
		wsServer.KeyFile = gate.KeyFile
		wsServer.NewAgent = func(conn *network.WSConn) network.Agent {
			a := &GateAgent{conn: conn, gate: gate, netType: WS}
			pid, err := gs.GetAgentActor(a)
			if err != nil {
				ERROR("New WS Gate Agent error: ", Inspect(err))
			}
			a.agentActor = pid
			return a
		}
	}

	var tcpServer *network.TCPServer
	if gate.TCPAddr != "" {
		tcpServer = new(network.TCPServer)
		tcpServer.Addr = gate.TCPAddr
		tcpServer.MaxConnNum = gate.MaxConnNum
		tcpServer.PendingWriteNum = gate.PendingWriteNum
		tcpServer.LenMsgLen = gate.LenMsgLen
		tcpServer.MaxMsgLen = gate.MaxMsgLen
		tcpServer.LittleEndian = gate.LittleEndian
		tcpServer.NewAgent = func(conn *network.TCPConn) network.Agent {
			a := &GateAgent{conn: conn, gate: gate, netType: TCP}
			ac, err := gs.GetAgentActor(a)
			if err != nil {
				ERROR("New TCP Gate Agent error: ", Inspect(err))
			}
			a.agentActor = ac
			return a
		}
	}

	if wsServer != nil {
		wsServer.Start()
	}
	if tcpServer != nil {
		tcpServer.Start()
	}

	gate.tcpServer = tcpServer
	gate.wsServer = wsServer
}

func (gate *Gate) OnDestroy() {
	if gate.wsServer != nil {
		gate.wsServer.Close()
	}
	if gate.tcpServer != nil {
		gate.tcpServer.Close()
	}
}