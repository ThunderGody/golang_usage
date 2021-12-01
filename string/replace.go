package main

import (
	"fmt"
	"strings"
)

func main() {
	var r = strings.NewReplacer("a", "A", "b", "B")
	strs := []string{"abc", "add", "bdd"}

	for _, s := range strs {
		sq := r.Replace(s)
		fmt.Println(sq)
	}
}
