package conc

import (
	"fmt"
	"sync"
	"time"
)

// utilize CSP-style concurrency or memory access synchronization
// aim for simplicity, use channels when possible, and treat goroutines like a free resource
type Counter struct {
	mu    sync.Mutex
	value int
}

// M:N scheduler maps M green threads to N OS threads
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func politeGoroutine() {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = 1 * time.Second
	greedyWorker := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(3 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}
		fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
	}
	politeWorker := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}
		fmt.Printf("Polite worker was able to execute %v work loops.\n", count)
	}
	wg.Add(2)
	go greedyWorker()
	go politeWorker()
	wg.Wait()
}
