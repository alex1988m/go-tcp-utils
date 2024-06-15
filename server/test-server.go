package server

import (
	"log"
	"net/http"
)

type TestServer struct {
	Addr string
	server *http.Server
}

func (r *TestServer) Start() error {
	r.server = &http.Server{
		Addr: r.Addr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.RemoteAddr, "connected")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("test server"))
		}),
	}
	if err := r.server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
