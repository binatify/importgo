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

	log.Fatal(http.ListenAndServe(":3000", nil))
}
