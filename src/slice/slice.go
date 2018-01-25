package main

import (
	"fmt"
)

func main() {
	// 使用 make 声明切片
	sliceA := make([]int, 5)
	for i := 0; i < 5; i++ {
		sliceA[i] = i
	}

	printHelper("sliceA", sliceA)

	// 切片可以复制
	sliceB := make([]int, 5)
	copy(sliceB, sliceA)
	printHelper("sliceB", sliceB)

	//切片长度可以变化
	sliceA = append(sliceA, 5)
	printHelper("new sliceA", sliceA)

	// 可以用数组来初始化切片
	arr := [5]int{1, 2, 3, 4, 5}
	sliceC := arr[0:3] // 左闭右开区间,包含：{1,2,3}
	printHelper("sliceC", sliceC)

	// 可以用切片来初始化切片
	sliceD := sliceC[:] // 区间[0,len(sliceC))
	printHelper("sliceD", sliceD)

	// 两个切片对应的底层数组是一致的，改动底层数组会同时影响到两个切片
	sliceD[0] = 5
	fmt.Println("after modify the underlying array:")
	printHelper("sliceC", sliceC)
	printHelper("sliceD", sliceD)
}

func printHelper(name string, slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%v[%v]: %v\n", name, i, slice[i])
	}

	fmt.Printf("len of %v: %v\n", name, len(slice))
	fmt.Printf("cap of %v: %v\n", name, cap(slice))

	fmt.Println()
}
