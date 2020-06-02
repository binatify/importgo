package main

import (
	"log"
	"net"
	"time"
)

var (
	defaultTimeout = time.Second * 30
)

func main() {
	// 尝试与 baidu.com:80 建立 tcp 连接
	c, err := net.Dial("tcp", "127.0.0.1:3000")

	if err != nil {
		log.Fatal(err)
	}

	defer c.Close() // 退出关闭连接，释放资源

	// 设置写超时时间
	_ = c.SetWriteDeadline(time.Now().Add(defaultTimeout))
	// 发送数据
	_, _ = c.Write([]byte("ping"))

	// 设置读超时时间
	_ = c.SetReadDeadline(time.Now().Add(defaultTimeout))
	//读取数据

	buff := make([]byte, 1024)
	n, _ := c.Read(buff)
	log.Printf("got response: %s\n", buff[:n])
}
