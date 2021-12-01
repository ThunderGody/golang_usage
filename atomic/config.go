package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	config := atomic.Value{}
	config.Store(22)

	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			if i == 0 {
				config.Store(23)
			}
			fmt.Println(config.Load())
		}(i)
	}

	wg.Wait()
}