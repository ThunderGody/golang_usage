package main

import "fmt"

func main() {
	startTime := "2020-11"

	endTime := "2021-11-01"

	startDay := startTime[0:10]
	endDay := endTime[0:10]

	fmt.Println(startDay)

	fmt.Println(endDay)
}
