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

	msg = <-ch // channel is closed
	fmt.Printf("Closed channels: %#v\n", msg)

	// receiving from a closed channel
	msg, ok := <-ch
	fmt.Printf("Closed channel: %#v, (%v)\n", msg, ok)

	testSleepSort := []int{4, 16, 3, 5, 8, 9, 9}
	fmt.Println(sleepSort(testSleepSort))
}

/*
Channel Semantics
- send and receive will block until opposite operations (*)
- sending through channels must be done inside of go routines
- Receiving from a closed channel will return an empty string
- Use the comma, ok to find out if you got zero because the value is missing or because that is the inputed value
- When using close(ch), after the channel has been read from, once, it can no longer be read from.
- Sending to a closed channel will panic
- Closing a closed channel will also panic, and there isnt any way to check if a channel is closed
- Send/receive to a nil channel will block forever
- You dont have to always close a channel
*/

/*
For every value "n" in values, spin a go routine that will
- sleep "n" milliseconds
- send "n" over a channel
- in the function body collect values from the channel to a slice and return it
*/

func sleepSort(values []int) []int {
	ch := make(chan int)

	for _, i := range values {
		i := i
		go func() {
			time.Sleep(time.Duration(i) * time.Millisecond)
			ch <- i
		}()
	}

	var collector []int
	for range values {
		i := <-ch
		collector = append(collector, i)
	}

	return collector
}
