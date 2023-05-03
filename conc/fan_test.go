package conc

import (
	"sync"
	"testing"
)

func fanIn(
	done <-chan interface{},
	channels ...<-chan interface{},
) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})
	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}
	// Select from all the channels
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}
	// Wait for all the reads to complete
	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()
	return multiplexedStream
}

func TestPrimesSlow(t *testing.T) {
	//rand := func() interface{} { return rand.Intn(50000000) }
	//done := make(chan interface{})
	//defer close(done)
	//start := time.Now()
	//randIntStream := toInt(done, repeatFn(done, rand))
	//fmt.Println("Primes:")
	//for prime := range take(done, primeFinder(done, randIntStream), 10) {
	//	fmt.Printf("\t%d\n", prime)
	//}
	//fmt.Printf("Search took: %v", time.Since(start))
}

func primeFinder(done chan interface{}, stream interface{}) <-chan interface{} {
	return nil
}
