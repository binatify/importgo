package main

import (
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		log.Fatalf("tcp listen with error: %s \n", err)
	}

	log.Printf("tcp server listen via %s", l.Addr().String())

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatalf("l.Accept(): %v", err)
		}

		log.Printf("tcp connect created \n")

		go handleConn(c)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	// 设置消息体缓冲区大小
	buff := make([]byte, 1024)

	for {
		n, err := c.Read(buff)

		if err != nil && err != io.EOF {
			log.Printf("c.Read(): %v", err)
		}
		if err != nil {
			return
		}

		log.Printf("got request: %s\n", string(buff[:n]))

		// 返回消息
		_, _ = c.Write([]byte("pong"))
	}
}
