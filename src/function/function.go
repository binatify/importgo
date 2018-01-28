package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func multi() (string, int) {
	return "age is: ", 18
}

func namedReturnValue() (name string, height int) {
	name = "禾木课堂"
	height = 180

	return
}

func Sum(nums ...int) int {
	fmt.Println("len of nums is : ", len(nums))
	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}

func sayHello(name string) {
	fmt.Println("Hello ", name)
}

func logger(f func(string), name string) {
	fmt.Println("start calling method sayHello")
	f(name)
	fmt.Println("end calling method sayHellog")
}

func sendValue(name string) {
	name = "hemuketang"
}

func sendAddress(name *string) {
	*name = "hemuketang"
}

func addInt(n int) func() int {
	i := 0
	return func() int {
		i += n
		return i
	}
}

func main() {
	// 单返回值函数
	res := plus(1, 2)
	fmt.Println(res)

	// 多返回值函数
	str, age := multi()
	fmt.Println(str)
	fmt.Println(age)

	// 命名返回值函数
	name, height := namedReturnValue()
	fmt.Println(name, " ", height)

	// 参数可变函数
	fmt.Println(Sum(1))
	fmt.Println(Sum(1, 2))
	fmt.Println(Sum(1, 2, 3))

	// 匿名函数
	func(name string) {
		fmt.Println(name)
	}("禾木课堂")
	logger(sayHello, "禾木课堂")

	// 传值和传引用
	str = "禾木课堂"
	fmt.Println("before calling sendValue, str : ", str)
	sendValue(str)
	fmt.Println("after calling sendValue, str : ", str)

	fmt.Println("before calling sendAddress, str : ", str)
	sendAddress(&str)
	fmt.Println("after calling sendAddress, str: ", str)

	// 闭包
	addOne := addInt(1)
	fmt.Println(addOne()) // 1
	fmt.Println(addOne()) // 2
	fmt.Println(addOne()) // 3

	addTwo := addInt(2)
	fmt.Println(addTwo()) // 2
	fmt.Println(addTwo()) // 4
	fmt.Println(addTwo()) // 6
}
