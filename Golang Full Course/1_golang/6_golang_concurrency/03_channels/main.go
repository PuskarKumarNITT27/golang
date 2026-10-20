package main

import (
	"fmt"
	"time"
)

func main() {
	// waitgroups helps you to wait

	// run work concurrently + collect results

	// pipe -> send values between goroutines

	// one goroutine sends:   ch <- value
	// another receives: value := <- ch

	// make(chan T)

	type User struct {
		ID   int
		Name string
	}

	// carries the User values
	ch := make(chan User)

	// worker goroutine
	go func() {
		// simulate slow work
		time.Sleep(200 * time.Millisecond)

		// Send: blocks until main receives
		// unbuffered channel , send + receives is a handshake
		ch <- User{ID: 100, Name: "Sangam"}

	}()

	fmt.Println("main: waiting to receive user...")

	u := <-ch

	fmt.Println("main: now go user", u, u.ID, u.Name)

}
