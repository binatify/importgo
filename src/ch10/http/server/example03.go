package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	handler := NewMyHandler()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,
			"Hello %s, current time is %s \n",
			r.URL.Query().Get("name"),
			time.Now().Format(time.RFC3339))
	})

	handler.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Pong\n")
	})

	log.Fatal(http.ListenAndServe(":3000", handler))
}

type MyHandler struct {
	mux    sync.RWMutex
	router map[string]HandFunc
}

type HandFunc func(w http.ResponseWriter, r *http.Request)

func NewMyHandler() *MyHandler {
	return &MyHandler{
		router: make(map[string]HandFunc, 8),
	}
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.RLock()
	hf, ok := h.router[r.URL.Path]
	h.mux.RUnlock()

	// if not found
	if !ok {
		http.NotFound(w, r)
		return
	}

	hf(w, r)
}

func (h *MyHandler) HandleFunc(pattern string, hf HandFunc) {
	h.mux.Lock()
	h.router[pattern] = hf
	h.mux.Unlock()
}
