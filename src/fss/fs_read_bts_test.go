package fss

import (
	"bytes"
	"fmt"
	"testing"
)

func TestReadOneByte(t *testing.T) {
	msg := "Lorem ipsum"

	reader := bytes.NewReader([]byte(msg))

	readByte, err := reader.ReadByte()
	if err != nil {
		fmt.Printf("Error reading byte %v \n", err)
	}

	fmt.Println(string(readByte))
}
