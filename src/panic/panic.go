package main

import (
	"errors"
	"fmt"
)

func deferPrint() (i int) {
	defer func() {
		fmt.Println(i)
		i = 4
	}()

	return 2
}

func panicTest() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("recover panic")
		}
	}()

	panic(errors.New("this is a panic"))
}

func main() {
	fmt.Println("before panic")

	panicTest()

	fmt.Println("after panic")
}
