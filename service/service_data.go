package service

import (

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
)

type IServiceData interface {
	Init(addr string, name string, typename string)
	GetName() string
	GetType() string
	GetAddress() string
	SetPID(pid *actor.PID)
	GetPID() *actor.PID
}

type ServiceData {
	Address string
	Name string
	TypeName string
	PID *actor.PID
}

func (d *ServiceData) Init(addr string, name string, typename string)  {
	d.Address = addr
	d.Name = name
	d.TypeName = typename
}

func (d *ServiceData) OnRun()  {
	
}

func (d *ServiceData) OnDestory()  {
	
}

func (d *ServiceData) GetType() string {
	return d.TypeName
}

func (d *ServiceData) GetAddress() string {
	return d.Address
}

func (d *ServiceData) GetPID() *actor.PID {
	return d.PID
}

func (d *ServiceData) SetPID(pid *actor.PID)  {
	d.PID = pid
}