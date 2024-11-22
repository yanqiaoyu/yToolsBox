package dspm

import "bytes"

func GetBytes(str string) []byte {
	buffer := bytes.NewBufferString(str)
	byteSlice := buffer.Bytes()
	return byteSlice
}
