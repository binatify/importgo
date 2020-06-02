package main

import (
	"log"
	"net"
)

func main() {
	c, err := net.Dial("udp", "127.0.0.1:3000")
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	// 发送消息
	_, err = c.Write([]byte("ping"))
	if err != nil {
		log.Fatal(err)
	}

	// MTU 一般小于 1500 字节
	buf := make([]byte, 1500)

	// 持续读数据
	for {
		n, _ := c.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(string(buf[:n]))
	}
}
