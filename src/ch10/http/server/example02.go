package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,
			"Hello %s, current time is %s \n",
			r.URL.Query().Get("name"),
			time.Now().Format(time.RFC3339))
	})

	server := &http.Server{
		Addr:              ":3000",
		Handler:           nil, // 具体的 handler，如果为空则使用默认的 http.DefaultServeMux
		TLSConfig:         nil, // TLS 相关的配置，如果不为空，那么默认
		ReadTimeout:       0,   // 整个 http request 的读超时时间，包括 header 和 body，如果为 0 表示永远不超时
		ReadHeaderTimeout: 0,   // http request header 读取超时时间设置，如果为 0 表示永远不超时
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
		// 其它字段 ..
	}

	log.Fatal(server.ListenAndServe())
}
