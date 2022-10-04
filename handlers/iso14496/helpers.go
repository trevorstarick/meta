package iso14496

import (
	"encoding/binary"
	"fmt"
	"io"
)

func GetAtom(r io.ReadSeeker) (int, []byte, error) {
	var length int32
	var name [4]byte

	if err := binary.Read(r, binary.BigEndian, &length); err != nil {
		return 0, nil, err
	}

	if err := binary.Read(r, binary.BigEndian, &name); err != nil {
		return 0, nil, err
	}

	return int(length), name[:], nil
}

func Skip(r io.ReadSeeker, l int) error {
	if l < 0 {
		return nil
	}

	_, err := r.Seek(int64(l), io.SeekCurrent)
	if err != nil {
		return err
	}

	return nil
}

func debug(r io.ReadSeeker, parent string, name []byte, length int) error {
	o, err := r.Seek(-8, io.SeekCurrent)
	if err != nil {
		return err
	}

	fmt.Printf("unknown %s atom: %s @ %d -> %d\n", parent, string(name), o, o+int64(length))

	_, err = r.Seek(o+int64(length), io.SeekStart)
	if err != nil {
		return err
	}

	return nil
}

var Debug = debug
