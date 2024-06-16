package server

import (
	"github/alex1988m/go-tcp-utils/handler"
	"log"
	"net"
)

type TCPServer struct {
	ServerAddr string
	Handler    handler.TCPHandler
	l          net.Listener
}

func (p *TCPServer) Start() error {
	if err := p.listen(); err != nil {
		return err
	}
	log.Printf("proxy server started on %s", p.ServerAddr)
	return nil
}

func (p *TCPServer) Stop() error {
	if err := p.l.Close(); err != nil {
		return err
	}
	log.Printf("proxy server stopped on %s", p.ServerAddr)
	return nil
}

func (p *TCPServer) Serve() error {
	for {
		conn, err := p.l.Accept()
		log.Printf("client connected: %s", conn.RemoteAddr().String())
		if err != nil {
			return err
		}
		go p.Handler.Handle(conn)
	}
}

func (p *TCPServer) listen() error {
	l, err := net.Listen("tcp", p.ServerAddr)
	if err != nil {
		return err
	}
	p.l = l
	return nil
}
