package corr

import (
	"fmt"
	"testing"
)

func TestInter1(t *testing.T) {
	valuesToAdd := [][]interface{}{{"a", "b"}, {"c", "d"}} //[["", ""],["", ""]]
	fmt.Println(valuesToAdd[0][0])
	fmt.Println(valuesToAdd[1][0])
}
