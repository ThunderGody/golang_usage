package main

import "fmt"

func main()  {
	s := "Hello,世界"
	fmt.Println(len(s))

	for i, r := range "Hello,世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}
