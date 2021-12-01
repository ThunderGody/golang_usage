package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int32
func add(wg *sync.WaitGroup) {
	defer wg.Done()

	atomic.AddInt32(&count, 1)
}


func main() {
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go add(&wg)
	}

	wg.Wait()

	fmt.Println(count)
}
