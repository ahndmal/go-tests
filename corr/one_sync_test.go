package corr

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSyncTwo(t *testing.T) {
	fmt.Println("Go Channel Tutorial")

	values := make(chan int)
	defer close(values)

	go calculate(values)

	value := <-values
	fmt.Println(value)
}

func calculate(values chan int) {
	val := rand.Intn(9)
	fmt.Println("Calculated Random Value: {}", val)
	values <- val
}
