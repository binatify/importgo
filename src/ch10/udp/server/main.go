package main

import (
	"log"
	"net"
	"time"
)

func main() {
	// 监听本地地址并等待接收包
	laddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:3000")
	c, err := net.ListenUDP("udp", laddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("udp server listen via : %s", c.LocalAddr())

	defer c.Close()

	// 缓存所有的客户端
	clients := make([]net.Addr, 0)

	// 分别向所有的客户端回包
	go func() {
		for {
			for _, addr := range clients {
				_, err := c.WriteTo([]byte("pong"), addr)
				if err != nil {
					log.Println(err)
				}
			}

			time.Sleep(5 * time.Second)
		}
	}()

	// 等待客户端的连接，并缓存客户端的地址
	for {
		buf := make([]byte, 256)
		n, addr, err := c.ReadFrom(buf)
		if err != nil {
			log.Println(err)
			continue
		}

		clients = append(clients, addr)
		log.Println(string(buf[0:n]))
		log.Println(addr.String(), "connecting...", len(clients), "connected")
	}
}
