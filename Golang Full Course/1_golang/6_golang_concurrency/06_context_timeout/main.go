package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// stuck ?
	// too long -> DB
	// context.Context
	// ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 450*time.Millisecond)

	defer cancel()

	go slowWork(ctx)

	// main waits until context ends
	<-ctx.Done()

	fmt.Println("main: context ended", ctx.Err())
	fmt.Println("main:exit")
}

// select + time.After
// context.withtimeout + ctx.Done()

func slowWork(ctx context.Context) {
	select {
	case <-time.After(700 * time.Millisecond):
		return

	case <-ctx.Done():
		return
	}
}
