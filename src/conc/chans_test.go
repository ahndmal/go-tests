package conc

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sync"
	"testing"
)

func TestBclockChann(t *testing.T) {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream
	}
	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}

func TestBuffChannels(t *testing.T) {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)
	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()
	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}
}

func TestUnblockGoroutines(t *testing.T) {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}
	fmt.Println("Unblocking goroutines...")
	close(begin)
	wg.Wait()
}

func TestClosingChans(t *testing.T) {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()
	for integer := range intStream {
		fmt.Printf("%v ", integer)
	}
}

func TestTwoChanValues(t *testing.T) {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"
	}()
	salutation, ok := <-stringStream
	fmt.Printf("(%v): %v", ok, salutation)
}

func TestChannExample(t *testing.T) {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"
	}()
	fmt.Println(<-stringStream)
}

func TestChannelsTypes(t *testing.T) {
	var dataStream chan interface{}
	dataStream = make(chan interface{})
	log.Println(dataStream)

	var readStream <-chan interface{}
	readStream2 := make(<-chan interface{})
	log.Println(readStream, readStream2)

	var writeStream chan<- interface{}
	writeStream2 := make(chan<- interface{})
	log.Println(writeStream, writeStream2)
	channelsUni := func() {
		var receiveChan <-chan interface{}
		var sendChan chan<- interface{}
		dataStream := make(chan interface{})
		// Valid statements:
		receiveChan = dataStream
		sendChan = dataStream
		log.Println(receiveChan, sendChan)
	}
	channelsUni()
}
