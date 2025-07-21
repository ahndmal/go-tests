package net

import (
	"io"
	"log"
	"net/http"
	"sync"
	"testing"
)

func TestUrlsAsync(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetData("http://example.com", &wg)
		}()
	}
	wg.Wait()
}

func GetData(url string, wg *sync.WaitGroup) {
	log.Println(" Getting data...")
	//defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(data))
}
