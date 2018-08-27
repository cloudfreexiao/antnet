package service

import (
	"fmt"
	"reflect"

	"github.com/cloudfreexiao/antnet/ecs"
	"github.com/cloudfreexiao/antnet/logger"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
)

type ActorContext actor.Context

type IService interface {
	IServiceData

	OnReceive(context actor.Context)
	OnInit()
	OnStart(as *ActorService)
	OnRun()
	OnDestory()
}

type ServiceRun struct {}
type MessageFunc func(context actor.Context)

// 服务代理
type ActorService struct {
	serviceIns IService
	router map[reflect.Type] MessageFunc
	ctx ecs.Context
}



func (s *ActorService) Receive(context actor.Context)  {
	switch msg := context.Message().(type) {
	case *actor.Started:
		DEBUG("++++++++actor service Started++++++++")
	case *actor.Stopping:
		DEBUG("++++++++actor service Stopping++++++++")
	case *actor.Restarting:
		DEBUG("++++++++actor service Restarting++++++++")
	case *actor.ServiceRun:
		DEBUG("++++++++actor service Restarting++++++++")
		s.serviceIns.OnRun()
	default:
		DEBUG("recv msg:", Inspect(msg))
		s.serviceIns.OnReceive(context.(ActorContext))
		fun := s.router[reflect.TypeOf(msg)]
		if fun != nil {
			fun(context.(ActorContext))
		} else {
			ERROR("recv msg but not found func:", Inspect(msg))
		}
	}
}

func (s *ActorService) RegisterMsg(t reflect.Type, fun MessageFunc)  {
	s.router[t] = fun
}


func StartService(s IService)  {
	ac := &ActorService{s, make(map[reflect.Type]MessageFunc), ecs.Context{}}
	props := actor.FromProducer(func ()  {
		actor.Actor {return ac}
	})

	if s.GetAddress() != "" {
		remote.Start(s.GetAddress())
	}

	pid, err := actor.SpawnNamed(props, s.GetName())
	if err == nil {
		s.SetPID(pid)
		s.OnStart(ac)
	} else {
		ERROR("Actor SpawnNamed error:", Inspect(err))
	}
}

func DestoryService(s *ActorService)  {
	s.serviceIns.OnDestory()
}
