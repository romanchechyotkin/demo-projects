package main

import (
	"fmt"
	"log"
	"math/rand"
	"mistakes/11_builder_pattern/server"
	"sync"
)

type ConfigBuilder struct {
	port *int
}

func (cb *ConfigBuilder) Port(p int) *ConfigBuilder {
	cb.port = &p
	return cb
}

func (cb *ConfigBuilder) Build() (*server.Config, error) {
	var cfg server.Config
	const DefaultPort = 6969
	if cb.port == nil {
		*cb.port = DefaultPort
	}

	if *cb.port < 0 {
		return nil, fmt.Errorf("negfative port: %d", *cb.port)
	}

	if *cb.port == 0 {
		*cb.port = rand.Intn(10000)
	}

	cfg.Port = *cb.port

	return &cfg, nil
}

func main() {
	wg := sync.WaitGroup{}

	builderF := &ConfigBuilder{}
	builderF.Port(6666)
	cfg, err := builderF.Build()
	if err != nil {
		log.Println(err)
	}
	s := server.New(cfg)

	builderF.Port(2288)
	cfg2, err := builderF.Build()
	if err != nil {
		log.Println(err)
	}

	s2 := server.New(cfg2)

	wg.Add(1)
	go s.Run()
	go s2.Run()

	wg.Wait()
}
