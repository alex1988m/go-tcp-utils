package handler

import (
	"io"
	"net"
	"os/exec"
)

type TCPControlHandler struct {
	Command []string
}

func (h *TCPControlHandler) Handle(client net.Conn) error {
	defer client.Close()
	cmd := exec.Command(h.Command[0], h.Command[1:]...)
	pr, pw := io.Pipe()
	cmd.Stdin = client
	cmd.Stdout = pw
	go io.Copy(client, pr)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
