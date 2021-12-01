package main

import "fmt"

func main() {
	z := make([]int, 0)
	for i := 0; i < 10; i++ {
		z = appendInt(z, i)
	}

	fmt.Println(z)
}
func appendInt(x []int, v int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen < cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2 * cap(x) {
			zcap = 2 * cap(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = v
	return z
}
