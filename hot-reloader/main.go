package main

import (
	"sync"
	"time"

	"hot_reloader/api"
	"hot_reloader/reloader"
	"hot_reloader/server"
)

func main() {
	var wg sync.WaitGroup

	a := api.New("https://example.com")
	srv := server.New(":8080", a)

	wg.Add(1)
	srv.Start()

	r := reloader.New("http", &reloader.ReloaderConfig{
		ConfigPath: "config.yaml",
		Timeout:    10 * time.Second,
	})
	r.Register(srv)
	r.Start()

	ar := reloader.New("api", &reloader.ReloaderConfig{
		ConfigPath: "api.yaml",
		Timeout:    10 * time.Second,
	})
	ar.Register(a)
	ar.Start()

	wg.Wait()
}
