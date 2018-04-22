package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(name string) {
	data, err := ioutil.ReadFile(name)
	errCheck(err)
	fmt.Println("data: ", string(data))
}

func main() {
	path := "test.txt"
	content := "hello"

	newFile, err := os.Create(path)
	errCheck(err)

	bufferWriter := bufio.NewWriter(newFile)

	for _, v := range content {
		n, err := bufferWriter.WriteString(string(v))
		errCheck(err)
		fmt.Println("n is : ", n)
	}

	readFile(path) // empty

	bufferWriter.Flush()

	readFile(path)
}
