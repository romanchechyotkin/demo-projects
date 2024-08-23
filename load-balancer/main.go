package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"sync"
)

// TODO implement to car-booking-service
// TODO Update project

type Server interface {
	IsAlive() bool
	Address() string
	Serve(w http.ResponseWriter, r *http.Request)
}

type Backend struct {
	URL          string
	Alive        bool
	mux          sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

func NewBackend(addr string) *Backend {
	u, err := url.Parse(addr)
	handleErr(err)

	return &Backend{
		URL:          addr,
		ReverseProxy: httputil.NewSingleHostReverseProxy(u),
	}
}

func (b *Backend) Address() string {
	return b.URL
}

func (b *Backend) IsAlive() bool {
	return true
}

func (b *Backend) Serve(w http.ResponseWriter, r *http.Request) {
	b.ReverseProxy.ServeHTTP(w, r)
}

type LoadBalancer struct {
	port            string
	backends        []Server
	robinRoundCount int
}

func NewLoadBalancer(port string, backends []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		backends:        backends,
		robinRoundCount: 0,
	}
}

func main() {
	backends := []Server{
		NewBackend("https://www.facebook.com"),
		NewBackend("https://www.youtube.com"),
		NewBackend("https://www.ya.ru"),
		NewBackend("https://www.instagram.com"),
	}

	lb := NewLoadBalancer(":8000", backends)
	handleRedirect := func(w http.ResponseWriter, r *http.Request) {
		lb.serveProxy(w, r)
	}
	http.HandleFunc("/", handleRedirect)
	log.Fatal(http.ListenAndServe(lb.port, nil))
}

func handleErr(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func (lb *LoadBalancer) getNextAvailableServer() Server {
	l := len(lb.backends)
	server := lb.backends[lb.robinRoundCount%l]
	for !server.IsAlive() {
		lb.robinRoundCount++
		server = lb.backends[lb.robinRoundCount%l]
	}
	lb.robinRoundCount++
	return server
}

func (lb *LoadBalancer) serveProxy(w http.ResponseWriter, r *http.Request) {
	targetServer := lb.getNextAvailableServer()
	log.Printf("requesting %s server", targetServer.Address())
	targetServer.Serve(w, r)
}
