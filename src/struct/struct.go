package main

import "fmt"

type Person struct {
	Age  int
	Name string
}

type Student struct {
	Person
}

func main() {
	stu := Student{
		Person{
			Age:  18,
			Name: "hemu",
		},
	}

	fmt.Println(stu)

	fmt.Println(stu.Age)
	fmt.Println(stu.Name)
}

// package main

// import "fmt"

// type Tree struct {
// 	value       int
// 	left, right *Tree
// }

// func main() {
// 	tree := Tree{
// 		value: 1,
// 		left: &Tree{
// 			value: 1,
// 			left:  nil,
// 			right: nil,
// 		},
// 		right: &Tree{
// 			value: 2,
// 			left:  nil,
// 			right: nil,
// 		},
// 	}

// 	fmt.Printf(">>> %#v\n", tree)
// }
