package main

import (
	"log"
	"net/http"
	"sync"
	"testing"
)

func ExampleWaitGroup() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.example.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			response, err := http.Get(url)
			if err != nil {
				log.Fatalln(err)
			}
			log.Println(response)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}
func TestSyncOne(t *testing.T) {
	ExampleWaitGroup()
	//
}
