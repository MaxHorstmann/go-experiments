package main

import (
	"context"
	"fmt"
	"time"
)

func doSomethingLongRunning(ctx context.Context) {
	for i := 0; i < 100; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("I've been cancelled!")
			return
		default:
			fmt.Println(i)
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go doSomethingLongRunning(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("Cancelling now, this is taking too long!")
	cancel()

	fmt.Println("Hit enter")
	fmt.Scanln()
}
