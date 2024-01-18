package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)

	someHttpRequest(ctx)

	// fmt.Println("context:\t", ctx)
	// fmt.Println("context err:\t", ctx.Err())
	// fmt.Printf("context type:\t%T\n", ctx)
	// fmt.Println("----------")

	// fmt.Println("context:\t", ctx)
	// fmt.Println("context err:\t", ctx.Err())
	// fmt.Printf("context type:\t%T\n", ctx)
	// fmt.Println("cancel:\t\t", cancel)
	// fmt.Printf("cancel type:\t%T\n", cancel)
	// fmt.Println("----------")

	defer cancel()

	// fmt.Println("context:\t", ctx)
	// fmt.Println("context err:\t", ctx.Err())
	// fmt.Printf("context type:\t%T\n", ctx)
	// fmt.Println("cancel:\t\t", cancel)
	// fmt.Printf("cancel type:\t%T\n", cancel)
	// fmt.Println("----------")
}

func someHttpRequest(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// The context is canceled
			fmt.Println("Worker: Context canceled")
			return
		default:
			// Do some work
			fmt.Println("Worker: Working...")
			time.Sleep(1 * time.Second)
		}
	}
}
