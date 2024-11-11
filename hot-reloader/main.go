package main

import (
	"sync"

	"hot_reloader/reloader"
	"hot_reloader/server"
)

func main() {
	var wg sync.WaitGroup

	srv := server.New(":8080")

	wg.Add(1)
	srv.Start()

	r := reloader.New("http", "config.yaml")
	r.Register(srv)

	r.Start()

	wg.Wait()
}
