package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func Async2() {
	fmt.Println("> start")

	go compute(10)
	go compute(10)

	//time.Sleep(3000)

	var input string
	fmt.Scanln(&input)

}

func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

//
//
//

func MyFunc(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("Finished executing Goroutine")
	wg.Done()
}

func main() {
	//ms := net.MySocket{}
	//ms.Cast()

	msg := "Lorem ipsum"
	reader := bytes.NewReader([]byte(msg))
	b1, _ := reader.ReadByte()
	b2, _ := reader.ReadByte()
	println(string(b1))
	println(string(b2))

}

func UrlPingerAsync2() {
	urls := []string{
		"https://us-central1-andmal-bot.cloudfunctions.net/gcp-java-perf-test", // java
		"https://us-central1-andmal-bot.cloudfunctions.net/node2",              // node
		"https://us-central1-andmal-bot.cloudfunctions.net/go-perf-test",       // node
		"https://us-central1-andmal-bot.cloudfunctions.net/python-perf-test",   // python
		"https://us-central1-andmal-bot.cloudfunctions.net/dotnet2",            // dotnet
	}

	//OneUrlReq(urls[0], 20)
	//OneUrlReq(urls[1], 20)
	//OneUrlReq(urls[2], 100)
	//OneUrlReq(urls[3], 100)
	//OneUrlReq(urls[4], 20)

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			OneUrlReq(url, 30)
			wg.Done()
		}(url)
	}
	wg.Wait() // block

	//time.Sleep(20000)

	//fmt.Println("Wait group init")
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go MyFunc(&wg)
	//wg.Wait()
	//fmt.Println("Finished Main")
}

func OneUrlReq(url string, times int) {
	//url := urls[0]
	for i := 0; i < times; i++ {
		http.Get(url)
		fmt.Printf(" -- Req %d for %s is DONE \n", i, url)
	}
}

func AllLinks(urls []string) {
	for _, url := range urls {
		fmt.Printf(" -- link is %s \n", url)
		for i := 0; i < 30; i++ {
			http.Get(url)
			fmt.Printf(" -- Req %d for %s is DONE \n", i, url)
		}
	}
}

//
//

func WaitEx2() {
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(50)

	for i := 0; i < 20; i++ {
		wg.Add(1)
		//for i := 0; i < 5; i++ {
		go func() {
			response, err := http.Get("https://example.com")
			if err != nil {
				log.Println(err)
			}
			fmt.Println(response.Body)
			defer wg.Done()
			//time.Sleep(200)
		}()
		//}
	}

	//wg.Wait()
}

func hClient() http.Client {
	return http.Client{Timeout: 15 * time.Second}
}
