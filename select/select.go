package main

import (
	"fmt"
	"time"

)

func main() {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		time.Sleep(10 * time.Millisecond)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(20 * time.Millisecond)
		ch2 <- 2
	}()

	/* Select works on several channels. It could be a send channel or a receive channel although, most times it is a receive channel*/
	select {
	case val := <-ch1:
		fmt.Println("ch1 : ", val)
	case val := <-ch2:
		fmt.Println("ch2 :", val)
	case <-time.After(5 * time.Millisecond):
		fmt.Println("Timeout")
	}
	/*
		Running this code would return ch1 all the time because its sleep time is smaller than channel 2's, but on reducing channel 2's sleep time,
		to less than channel 1's, it returns channel 2's value.
		An empty select with nothing inside returns a block forever
	*/
}
