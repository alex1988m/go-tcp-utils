package handler

import (
	"io"
	"log"
	"net"
)

type TCPProxyHandler struct {
	TargetAddr string
}

func (h *TCPProxyHandler) Handle(client net.Conn) error {
	defer client.Close()
	log.Printf("try to connect to %s", h.TargetAddr)
	target, err := net.Dial("tcp", h.TargetAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("target connected: %s", target.RemoteAddr().String())
	errCh := make(chan error, 1)
	defer target.Close()
	go func() {
		_, err := io.Copy(target, client)
		errCh <- err
	}()
	go func() {
		_, err := io.Copy(client, target)
		errCh <- err
	}()
	for i := 0; i < 2; i++ {
		err := <-errCh
		if err != nil {
			return err
		}
	}
	log.Println("success: both connections closed")
	return nil
}
