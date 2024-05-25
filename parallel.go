package main


type Work struct {
    x, y, z int
}


func main() {

	timerChan := make(chan time.Time)
	go func() {
    	time.Sleep(deltaT)
    	timerChan <- time.Now() // send time on timerChan
	}()
	// Do something else; when ready, receive.
	// Receive will block until timerChan delivers.
	// Value sent is other goroutine's completion time.
	completedAt := <-timerChan

	// select
	select {
	case v := <-ch1:
	    fmt.Println("channel 1 sends", v)
	case v := <-ch2:
	    fmt.Println("channel 2 sends", v)
	default: // optional
	    fmt.Println("neither channel was ready")
	}

	//
	
	
	
}

func daemon() {
	go func() { // copy input to output
	    for val := range input {
	        output <- val
	    }
	}()
}
