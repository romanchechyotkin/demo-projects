package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"hot_reloader/api"
	"hot_reloader/config"
)

type Server struct {
	httpsrv *http.Server
	api     *api.Api
}

func New(address string, a *api.Api) *Server {
	srv := &Server{
		api: a,
	}

	srv.registerRoutes()

	srv.httpsrv = &http.Server{
		Addr: address,
	}

	return srv
}

func (s *Server) registerRoutes() {
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		code := s.api.MakeRequest()
		log.Println("got request", "method", r.Method, "url", r.URL)
		fmt.Fprintf(w, "%d\n", code)
	})
}

func (s *Server) Start() {
	go func() {
		log.Println("Starting HTTP server on", s.httpsrv.Addr)
		if err := s.httpsrv.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Println("HTTP server closed:", s.httpsrv.Addr)
			} else {
				log.Println("Failed to start HTTP server:", err)
			}
		}
	}()
}

func (s *Server) Stop() {
	// s.mu.Lock()
	// defer s.mu.Unlock()

	log.Println("stopping http server", s.httpsrv.Addr)
	if err := s.httpsrv.Shutdown(context.Background()); err != nil {
		log.Println("failed to shutdown http server", err)
	}
}

func (s *Server) Reload(cfg any) {
	s.Stop()

	newAddr := cfg.(*config.Config).Address
	log.Println("Reloading HTTP server with new address:", newAddr)

	s.httpsrv = &http.Server{
		Addr:    newAddr,
		Handler: s.httpsrv.Handler,
	}

	s.Start()
}
