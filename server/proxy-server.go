package server

import (
	"io"
	"log"
	"net"
)

type ProxyServer struct {
	ProxyAddr  string
	TargetAddr string
	l          net.Listener
}

func (p *ProxyServer) Start() error {
	err := p.listen()
	log.Printf("proxy server started on %s", p.ProxyAddr)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProxyServer) Stop() {
	p.l.Close()
}

func (p *ProxyServer) Serve() error {
	for {
		conn, err := p.l.Accept()
		log.Printf("client connected: %s", conn.RemoteAddr().String())
		if err != nil {
			return err
		}
		go p.handle(conn)
	}
}

func (p *ProxyServer) handle(client net.Conn) {
	defer client.Close()
	log.Printf("try to connect to %s", p.TargetAddr)
	target, err := net.Dial("tcp", p.TargetAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("target connected: %s", target.RemoteAddr().String())
	defer target.Close()
	go func() {
		_, err := io.Copy(target, client)
		if err != nil {
			log.Fatalln(err)
		}
	}()
	_, copyErr := io.Copy(client, target)
	if copyErr != nil {
		log.Fatalln(copyErr)
	}
}

func (p *ProxyServer) listen() error {
	l, err := net.Listen("tcp", p.ProxyAddr)
	if err != nil {
		return err
	}
	p.l = l
	return nil
}
