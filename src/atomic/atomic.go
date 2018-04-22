package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var (
		total int32 = 0
	)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				atomic.AddInt32(&total, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	// if atomic.CompareAndSwapInt32(&total, 0, 100) {
	// 	fmt.Println("compare and swap successfully")
	// } else {
	// 	fmt.Println("compare and swap failed")
	// }

	// fmt.Println("old value is : ", atomic.SwapInt32(&total, 100))

	fmt.Println("before atomic.Store, the value is : ", total)

	atomic.StoreInt32(&total, 100)

	fmt.Println("finally total is : ", atomic.LoadInt32(&total))
}
