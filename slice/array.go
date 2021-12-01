package main

import "fmt"

func f(a [3]int) {
	a[2] = 10
	fmt.Printf("%p, %v\n", &a, a)

}

func main() {
	a := [3]int{1, 2, 3}
	b := a
	a[2] = 4

	fmt.Printf("%p, %v\n", &a, a)
	fmt.Printf("%p, %v\n", &b, b)
	f(a)

	c := struct {
		s [3]int
	}{
		s: a,
	}

	d := c

	c.s[2] = 30
	d.s[2] = 20

	fmt.Printf("%p, %v\n", &a, a)

	fmt.Printf("%p, %v\n", &c, c)
	fmt.Printf("%p, %v\n", &d, d)
}
