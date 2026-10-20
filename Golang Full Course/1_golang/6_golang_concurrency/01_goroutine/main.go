package main

import (
	"fmt"
	"time"
)

func main() {

	// concepts ->
	// doing multiple tasks in a way that they can overlap in time
	// overlap -> while task A (network, db), Task b can run

	// overlapping -> concurrency
	// executing multiple tasks at the exact same time
	// golang code can be concurrent without being parallel

	// IO- bound work -> slow -> because of waiting
	// http req
	// db queries
	// calling microservices
	// waiting for timers/ retries

	// backend -> multiple service calls -> combine all results

	// not to use concurrency ?
	// simple

	// Goroutine
	// concurrent execution of a function
	// go someFunc() -> starts this concurrently

	start := time.Now()
	// A
	go func() {
		time.Sleep(300 * time.Millisecond)

		fmt.Println("goroutine A: finished simulated API at", time.Since(start))

	}()

	// GoRoutine B
	go func() {
		time.Sleep(150 * time.Millisecond)

		fmt.Println("goroutine B: finished simulated API at", time.Since(start))
	}()

	// main -> no waiting
	fmt.Println("main: started two go routine at", time.Since(start))

	// small work -> any logic
	fmt.Println("main: doing step 1", time.Since(start))
	time.Sleep(100 * time.Millisecond)

	fmt.Println("main: doing step 2", time.Since(start))
	time.Sleep(100 * time.Millisecond)

	fmt.Println("main: doing step 3", time.Since(start))

	//temp sleep time
	time.Sleep(500 * time.Millisecond)

	fmt.Println("main: exiting at", time.Since(start))
}
