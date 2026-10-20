package main

import (
	"fmt"
	"time"
)

func main() {
	// unbuffered -> handshake

	// buffered channel -> capacity -> buffer size
	// make(chan int, 3) ->

	// 2 means -> this channel can store 2 jobs without a receiver
	jobs := make(chan string, 2)

	go func() {
		fmt.Println("producer: sending job-1")
		jobs <- "job-1" // goes into the buffer (buffer has space)

		fmt.Println("producer: sending job-2")
		jobs <- "job-2" // goes into the buffer (buffer still has space here)

		fmt.Println("producer: sending job-3 but this will wait until consumer reads")
		jobs <- "job-3"

		fmt.Println("producer: sent all jobs")
		close(jobs) // close the sender -> no more jobs, all r done

	}()

	for job := range jobs {
		fmt.Println("consumer got", job)

		time.Sleep(300 * time.Millisecond)

		fmt.Println("consumer finished", job)
	}

	fmt.Println("main: all jobs are completed")
}
