package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"testing"
)

var urls = []string{
	"https://google.com",
	"https://example.com",
}

func LinksStatusTest(t *testing.T) {
	fmt.Println("WaitGroup init")
	http.HandleFunc("/", fetchStatus)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func fetchStatus(w http.ResponseWriter, request *http.Request) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(w, "%+v\n", err)
			}
			fmt.Fprintf(w, "%+v\n", resp)
			wg.Done()
		}(url)
	}
	wg.Wait() // block
}
