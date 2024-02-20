package main

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{"https://google.com", "https://apple.com",
		"https://doesntexist.biz"}

	var wg sync.WaitGroup // is essentially a counter that decrements when goroutine is done
	wg.Add(len(urls))
	for _, url := range urls {
		//wg =1 is another way this could be declared instead of using the add function
		url := url // shadowing
		go func() {
			defer wg.Done() // decrements the wait group by 1
			siteTime(url)
		}()
	}
	wg.Wait()
}

func siteTime(url string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error: %s -> %s", url, err)
		return
	}
	defer resp.Body.Close()

	if _, err := io.Copy(io.Discard, resp.Body); err != nil { // since i dont want to copy the body anywhere in particular
		log.Printf("Error: %s -> %s", url, err)
	}

	duration := time.Since(start)
	log.Printf("INFO: %s -> %v", url, duration)
}
