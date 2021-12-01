package main

import (
	"context"
	"fmt"
	"time"
)

type otherContext struct {
	context.Context
}

func main() {
	ctxa, cancel := context.WithCancel(context.Background())
	go worker(ctxa, "work1")

	tm := time.Now().Add(3 * time.Second)
	ctxb, _ := context.WithDeadline(ctxa, tm)

	go worker(ctxb, "work2")

	oc := otherContext{ctxb}
	ctxc := context.WithValue(oc, "key", "pass from main")
	go workWithValue(ctxc, "work3")

	time.Sleep(10 * time.Second)

	fmt.Println("-------- cancel--------------")
	cancel()

	time.Sleep(5 * time.Second)
	fmt.Println("main stop")
}

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			fmt.Printf("%s is running \n", name)
			time.Sleep(1 * time.Second)
		}
	}
}

func workWithValue(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			value := ctx.Value("key").(string)
			fmt.Printf("%s is running value=%s \n", name, value)
			time.Sleep(1 * time.Second)
		}
	}
}


