package main

import (
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

func main() {

	Async2()
}

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
