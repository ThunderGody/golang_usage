package main

import "fmt"

func main() {
	const World = "world"
	var hello = "hello"

	var helloWorld = hello + " " + World
	fmt.Println(helloWorld)

	fmt.Println(hello == "hello")
	fmt.Println(hello > helloWorld)
}
