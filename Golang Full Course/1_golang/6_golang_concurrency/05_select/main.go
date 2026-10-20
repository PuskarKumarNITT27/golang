package main

import (
	"fmt"
	"time"
)

func main() {
	// select -> switch for ur channels
	// it lets a goroutine wait on multiple channels operations at once

	// give me res but dont hang forever
	// user involvements -> user cancels or user is doing any action -> stop
	// timeout -> stop

	resultCh := make(chan string)

	// worker goroutine

	go func() {
		// simulate slow work / consider u r doing a network call
		time.Sleep(40 * time.Millisecond)

		resultCh <- "worker: success"

	}()

	// timeout channel
	timeoutCh := time.After(250 * time.Millisecond)

	select {
	case res := <-resultCh:
		fmt.Println("main: go result", res)

	case <-timeoutCh:
		fmt.Println("main: timeout happened, stop waiting")
	}

	fmt.Println("main: work is now done")
}
