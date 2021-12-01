package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter = struct{
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}

	// read
	counter.RLock()
	n := counter.m["some_key"]
	counter.RUnlock()
	fmt.Println("some_key:", n)

	// write
	counter.Lock()
	counter.m["some_key"]++
	counter.Unlock()
}
