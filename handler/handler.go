package handler

import "net"

type TCPHandler interface {
	Handle(client net.Conn) error
}
