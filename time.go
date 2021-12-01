package main

import (
	"fmt"
	"strings"
)

func main() {
	timeStr := "2021-01-09"
	fmt.Println(BuildSendTime(timeStr))
	var err error
	if err != nil {
		fmt.Println("err is not nil")
	}

}

func BuildSendTime(timeStr string) string {
	timeStr = strings.TrimSpace(timeStr)
	if len(timeStr) < 10 {
		return ""
	}

	subTimeStr := timeStr[0:10]
	timeSlice := strings.Split(subTimeStr, "-")
	if len(timeSlice) < 3 {
		return ""
	}
	sendTimeStr := fmt.Sprintf("预计%s月%s日发货", timeSlice[1], timeSlice[2])
	return sendTimeStr
}