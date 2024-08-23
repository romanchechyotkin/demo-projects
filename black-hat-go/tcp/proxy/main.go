package main

import (
	"io"
	"log"
	"net"
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

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "79.133.182.160:3000")
	if err != nil {
		log.Println(err)
		return
	}
	defer dst.Close()

	go func() {
		// Копируем вывод источника в получателя
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	// Копирование вывода получателя обратно в источник
	if _, err = io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}

}
