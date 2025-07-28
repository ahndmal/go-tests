package corr

import (
	"fmt"
	"testing"
	"time"
)

func TestPointers(t *testing.T) {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

func TestFixedZone(t *testing.T) {
	loc := time.FixedZone("UTC-8", -8*60*60)
	tm := time.Date(2009, time.November, 10, 23, 0, 0, 0, loc)
	fmt.Println("The time is:", tm.Format(time.RFC822))
}
