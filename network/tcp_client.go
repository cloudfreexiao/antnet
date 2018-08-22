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

func (tcpClient *TCPClient) Start()  {
	DEBUG("++++TCPClient Init++++", tcpClient.ConnNum)
	tcpClient.init()
	DEBUG("++++TCPClient Will Conn++++", tcpClient.ConnNum)

	for i := 0; i< tcpClient.ConnNum; i++ {
		tcpClient.wg.Add(1)
		go tcpClient.connect()
	}
}

func (tcpClient *TCPClient) init()  {
	tcpClient.Lock()
	defer tcpClient.Unlock()

	if tcpClient.ConnNum <= 0 {
		tcpClient.ConnNum = 1
		DEBUG("invalid ConnNum, reset to ", tcpClient.ConnNum )
	}

	if tcpClient.ConnectInterval <= 0 {
		tcpClient.ConnectInterval = 3 * time.Second
		DEBUG("invalid ConnectInterval, reset to ", tcpClient.ConnectInterval)
	}

	if tcpClient.PendingWriteNum <= 0 {
		tcpClient.PendingWriteNum = 100
		DEBUG("invalid PendingWriteNum, reset to ", tcpClient.PendingWriteNum)
	}

	if tcpClient.NewAgent == nil {
		ERROR("TCPClient NewAgent Must Not Be Nil")
	}

	if tcpClient.conns != nil {
		ERROR("TCPClient is Running")
	}

	tcpClient.conns = make(ConnSet)
	tcpClient.closeFlag = false

	// msg parser
	msgParser := NewMsgParser()
	msgParser.SetMsgLen(tcpClient.LenMsgLen, tcpClient.MinMsgLen, tcpClient.MaxMsgLen)
	msgParser.SetByteOrder(tcpClient.LittleEndian)
	tcpClient.msgParser = msgParser
}

func (tcpClient *TCPClient) connect()  {
	
}