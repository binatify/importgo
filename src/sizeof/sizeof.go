package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var (
		a      int   = 64
		a32bit int32 = 64
		a64bit int64 = 64
	)

	// 与操作系统位数有关，32 位操作系统输出为 4， 64位输出为 8
	fmt.Println(unsafe.Sizeof(a))

	// 输出为 4
	fmt.Println(unsafe.Sizeof(a32bit))

	// 输出为 8
	fmt.Println(unsafe.Sizeof(a64bit))
}
