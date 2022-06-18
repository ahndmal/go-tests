package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(50)

	for i := 0; i < 20; i++ {
		wg.Add(1)
		sp := fmt.Sprintf("DEV%d", i)
		for i := 0; i < 5; i++ {
			go func() {
				fmt.Printf("%sspace created", sp)
				defer wg.Done()
				time.Sleep(200)
			}()
		}
	}

	wg.Wait()
}
