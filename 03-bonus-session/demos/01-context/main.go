package main

import (
	"context"
	"fmt"
	"time"
)

func slowLog(ctx context.Context, msg string) {
	fmt.Println("goroutine starting select")
	fmt.Println(ctx.Value("myKey"))

	select {
	case <-time.After(time.Second * 5):
		fmt.Println(msg)
	case <-ctx.Done():
		fmt.Println("context done msg")
	}

	fmt.Println("goroutine exits")
}

func main() {
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second*3)

	ctxWithValue := context.WithValue(ctxWithTimeout, "myKey", "myValue")
	go slowLog(ctxWithValue, "hello world")

	time.Sleep(time.Millisecond * 1000)
	cancel()

	time.Sleep(time.Millisecond * 3100)
	fmt.Println("main exits")
}
