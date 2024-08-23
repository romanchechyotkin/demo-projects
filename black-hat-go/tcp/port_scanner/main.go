package main

import (
	"fmt"
	"log"
	"net"
	"slices"
)

func scanPort(host string, port int) error {
	address := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		return err
	}
	defer conn.Close()

	log.Printf("connected to %s", conn.RemoteAddr())
	return nil
}

func worker(ports chan int, results chan int) {
	for port := range ports {
		if err := scanPort("scanme.nmap.org", port); err != nil {
			results <- 0
			continue
		}

		results <- port
	}
}

func main() {
	portsChan := make(chan int, 100)
	resultsChan := make(chan int)
	var openPorts []int

	for i := 1; i <= cap(portsChan); i++ {
		go worker(portsChan, resultsChan)
	}

	go func() {
		for i := 0; i < 1024; i++ {
			if i == 25 {
				portsChan <- 0
			}
			portsChan <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-resultsChan
		if port > 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(portsChan)
	close(resultsChan)
	slices.Sort(openPorts)

	for _, port := range openPorts {
		log.Println("open port", port)
	}

}
