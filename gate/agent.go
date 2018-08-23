package gate

import (
	"net"
	"reflect"

	"github.com/cloudfreexiao/antnet/logger"
	"github.com/cloudfreexiao/antnet/network"
	"github.com/cloudfreexiao/proto/messages"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type NetType = byte
const (
	TCP NetType = 0
	WS  NetType = 1
)

type Agent interface {
	WriteMsg(msg []byte)
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
	UserData() interface{}
	SetUserData(data interface{})
	SetDead()
	GetNetType() NetType
}

type GateAgent struct {
	conn network.Conn
	gate *Gate
	agentActor *actor.PID
	userData interface {}
	dead bool
	netType NetType
}

func (a *GateAgent) GetNetType() NetType {
	return a.netType
}

func (a *GateAgent) SetDead()  {
	a.dead = true
}

func (a *GateAgent) LocalAddr() net.Addr {
	return a.conn.LocalAddr()
}

func (a *GateAgent) RemoteAddr() net.Addr {
	return a.conn.RemoteAddr()
}

func (a *GateAgent) Close() {
	a.conn.Close()
}

func (a *GateAgent) Destroy() {
	a.conn.Destroy()
}

func (a *GateAgent) UserData() interface{} {
	return a.userData
}

func (a *GateAgent) SetUserData(data interface{}) {
	a.userData = data
}

func (a *GateAgent) Run() {
	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			DEBUG("read message: ", Inspect(err))
			break
		}

		if a.gate.Processor != nil {
			msg, err := a.gate.Processor.Unmarshal(data)
			if err != nil {
				DEBUG("unmarshal message error: ", Inspect(err) )
				break
			}
			err = a.gate.Processor.Route(msg, a)
			if err != nil {
				DEBUG("route message error: ", Inspect(err) )
				break
			}

		} else {
			//todo:not safe
			a.agentActor.Tell(&messages.ReceviceClientMsg{data})
		}
	}
}

func (a *GateAgent) OnClose() {
	if a.agentActor != nil && !a.dead {
		a.agentActor.Tell(&messages.ClientDisconnect{})
	}

}

func (a *GateAgent) WriteMsg(data []byte) {
	err := a.conn.WriteMsg(data)
	if err != nil {
		ERROR("write message: ",Inspect(reflect.TypeOf(data)), " error:", Inspect(err) )
	}
}