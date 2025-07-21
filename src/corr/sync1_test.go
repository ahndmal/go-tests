package corr

import (
	"fmt"
	"testing"
	"time"
)

func TestAsync2(t *testing.T) {
	fmt.Println("> start")

	go compute2(10)
	go compute2(10)

	time.Sleep(3000)

	var input string
	fmt.Scanln(&input)

}

func compute2(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
