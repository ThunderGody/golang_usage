package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type T struct {
	x int
}

var pT *T

func main() {
	var unsagePPT = (*unsafe.Pointer)(unsafe.Pointer(&pT))
	var ta, tb = T{1}, T{2}
	atomic.StorePointer(unsagePPT, unsafe.Pointer(&ta))
	fmt.Println(pT)

	pa1 := (*T) (atomic.LoadPointer(unsagePPT))
	fmt.Println(pa1 == &ta)

	pa2 := atomic.SwapPointer(unsagePPT, unsafe.Pointer(&tb))
	fmt.Println((*T) (pa2) == &ta)
	fmt.Println(pT)

	b := atomic.CompareAndSwapPointer(unsagePPT, pa2, unsafe.Pointer(&tb))
	fmt.Println(b)

	b = atomic.CompareAndSwapPointer(unsagePPT, unsafe.Pointer(&tb), pa2)
	fmt.Println(b)
}