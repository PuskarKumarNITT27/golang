package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// waitgroup -> counter lets main -> wait -> set of goroutines
	// add how many routines u have

	var wg sync.WaitGroup

	// i will wait for 3 routine
	wg.Add(3)

	go func() {

		defer wg.Done()

		fmt.Println("task 1")
		time.Sleep(250 * time.Millisecond)
		fmt.Println("task 1 is now donw")

	}()

	go func() {

		defer wg.Done()

		fmt.Println("task 2")
		time.Sleep(150 * time.Millisecond)
		fmt.Println("task 2 is now donw")

	}()

	go func() {

		defer wg.Done()

		fmt.Println("task 3")
		time.Sleep(199 * time.Millisecond)
		fmt.Println("task 3 is now donw")

	}()

	fmt.Println("main: waiting for tasks to finish")

	wg.Wait()

	fmt.Println("main: all tasks r finished, good job!")
}
