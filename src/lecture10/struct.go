package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Student struct {
	Person
}

func main() {
	per := Person{
		Name: "name",
		Age:  18,
	}

	stu := Student{Person: per}

	fmt.Println("stu.Name: ", stu.Name)
	fmt.Println("stu.Age: ", stu.Age)
}
