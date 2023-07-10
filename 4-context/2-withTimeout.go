package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	HttpHandler()
}

func HttpHandler() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	deal(ctx, cancel)
}

func deal(ctx context.Context, cancel context.CancelFunc) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Printf("deal time is %d\n", i)
			cancel()
		}
	}
}
