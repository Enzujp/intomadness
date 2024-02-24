package main

import (
	"fmt"
	"time"
)

// on using channels

func main() {
	fmt.Println("This is the first")
	go fmt.Println("This is the next")

	for i := 0; i < 3; i++ {
		i := i // here, i shadows the value of i 
		go func ()  {
			fmt.Println(i)
		}()
	}
	time.Sleep(10 * time.Millisecond)

}
/* 
The main function runs faster than the goroutine, because the goroutine is sectionized
to another portion of the cpu and this takes sometime.
Using the time.sleep function to delay the main, allows the goroutine to pick up some pace
*/