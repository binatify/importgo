package main

import "fmt"

func worker(ch chan string) {
	fmt.Println("get into worker...")
	ch <- "str"
}

func main() {
	ch := make(chan string)

	go worker(ch)

	<-ch
}
