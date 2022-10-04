package handlers

import (
	"encoding/binary"
	"io"
)

func atomToInt(b []byte) int {
	return int(b[0])<<24 + int(b[1])<<16 + int(b[2])<<8 + int(b[3])
}

func read(r io.Reader, v interface{}) error {
	return binary.Read(r, binary.BigEndian, v)
}
