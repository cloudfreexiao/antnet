package network

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"github.com/cloudfreexiao/antnet/logger"
)

type WSClient struct {
	sync.Mutex
	Addr string
	ConnNum int
	
}