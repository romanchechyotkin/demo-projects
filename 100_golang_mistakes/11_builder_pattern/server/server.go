package server

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct {
	Port int
}

type Server struct {
	base *http.Server
}

func New(cfg *Config) *Server {
	return &Server{base: &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.Port),
	}}
}

func (srv *Server) Run() {
	log.Println("running in port", srv.base.Addr)
	err := srv.base.ListenAndServe()
	if err != nil {
		return
	}
}
