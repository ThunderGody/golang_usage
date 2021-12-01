package main

import "fmt"

func main() {
	s := make([]int, 2)
	for i := 0; i < 2; i++ {
		s = append(s, i)
	}
	fmt.Println(s)

	s1 := make([]int, 0, 2)
	for i := 0; i < 2; i++ {
		s1 = append(s1, i)
	}
	fmt.Println(s1)
}
