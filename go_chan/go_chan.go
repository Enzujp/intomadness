package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("goroutine")
	fmt.Println("main")

	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i)
	}
	// another way would be to shadow the value of i,
	for i := 0; i < 4; i++ {
		i := i
		go func() {
			fmt.Println(i)
		}()
	}
	/* This would only print out the "main", because the main runtime does not wait for goroutine
	adding a sleep time function however gives goroutine time to sychronize and run*/
	time.Sleep(10 * time.Millisecond) /* this is really bad and not advisable however
	you cannot terminate a goroutine, unlike a process; once it starts, it is on its own and it goes til the end.
	Defer does not work on goroutines.
	*/

	ch := make(chan string)
	go func() {
		ch <- "hi" // send opertation
	}()

	msg := <-ch // receive
	fmt.Println(msg)

	go func() {
		for i := 0; i < 3; i++ {
			msg := fmt.Sprintf("Message: #%d", i+1)
			ch <- msg
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println("Got: ", msg)
	} /*
		This runs but it throws a deadlock error too, and this is because when you loop over a channel, go has no idea how many messages are coming
		- through, unlike with slices or arrays.
		This can be remedied by closing the channel, after sending in messages.	
	*/
}

/*
Channel Semantics
- send and receive will block until opposite operations (*)
- sending through channels must be done inside of go routines
*/
