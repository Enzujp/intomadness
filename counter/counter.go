package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex // customary to put Mutexes above what they guard
	count := 0

	const n = 10
	var Wg sync.WaitGroup
	Wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer Wg.Done()
			for j := 0; j < 10000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	Wg.Wait()
	fmt.Println(count)
}

// The Race detector is used in testing.
