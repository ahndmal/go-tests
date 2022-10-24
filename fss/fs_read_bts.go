package fss

import "bytes"

func ReadFileOne() (byte, error) {
	msg := "Lorem ipsum"
	reader := bytes.NewReader([]byte(msg))
	readByte, err := reader.ReadByte()
	if err != nil {
		return 0, err
	}
	return readByte, nil
}
