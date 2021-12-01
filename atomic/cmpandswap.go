package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var n int64 = 123
	var old = atomic.SwapInt64(&n, 789)
	fmt.Println(n, old)
	swapped := atomic.CompareAndSwapInt64(&n, 123, 456)
	fmt.Println(swapped)
	fmt.Println(n)

	swapped = atomic.CompareAndSwapInt64(&n, 789, 456)
	fmt.Println(swapped)
	fmt.Println(n)

}
