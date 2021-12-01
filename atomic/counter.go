package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var a int32

	fmt.Println("a:", a)

	newA := atomic.AddInt32(&a, 3)
	fmt.Println("newA:", newA)

	newA = atomic.AddInt32(&a, -2)
	fmt.Println("newA:", newA)


	var b int32
	fmt.Println("b:", b)

	atomic.CompareAndSwapInt32(&b, 0, 3)
	fmt.Println("b:", b)

	var c int32
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tmp := atomic.LoadInt32(&c)
			if !atomic.CompareAndSwapInt32(&c, tmp, (tmp + 1)) {
				fmt.Println("c update failure")
			}
		}()
	}

	wg.Wait()

	fmt.Println("c: ", c)

	var d int32
	fmt.Println("d:", d)

	atomic.StoreInt32(&d, 666)
	fmt.Println("d: ", d)

	var e int32
	wg2 := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			tmp := atomic.LoadInt32(&e)
			old := atomic.SwapInt32(&e, (tmp + 1))
			fmt.Println("e old:", old)
		}()
	}

	wg2.Wait()

	fmt.Println("e: ", e)
}