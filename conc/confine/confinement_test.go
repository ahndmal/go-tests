package confine

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestChannCloseAfterDuration(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}

func TestOrs(t *testing.T) {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}
		orDone := make(chan interface{})
		go func() {
			defer close(orDone)
			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()
		return orDone
	}
}

func TestChildGoroutCancel(t *testing.T) {
	doWork := func(
		done <-chan interface{},
		strings <-chan string,
	) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					// Do something interesting
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}
	done := make(chan interface{})
	terminated := doWork(done, nil)
	go func() {
		// Cancel the operation after 1 second.
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()
	<-terminated
	fmt.Println("Done.")
}

func TestGoroutineLeak(t *testing.T) {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				// Do something interesting
				fmt.Println(s)
			}
		}()
		return completed
	}
	doWork(nil)
	// Perhaps more work is done here
	fmt.Println("Done.")
}

func TestDataStrNotConcurSafe(t *testing.T) {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()
		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}
	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])
	wg.Wait()
}

func TestChanOwner(t *testing.T) {
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}
	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}
	results := chanOwner()
	consumer(results)
}

func TestAdHocConfinement(t *testing.T) {
	data := make([]int, 4)

	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}
