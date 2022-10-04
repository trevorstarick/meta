package iso14496

import (
	"encoding/binary"
	"io"
)

type VMHD struct {
	Version      byte
	Flags        [3]byte
	GraphicsMode int16
	OpColor      [3]int16
}

func (v *VMHD) Parse(r io.ReadSeeker, l int) error {
	var buf [4]byte

	if err := binary.Read(r, binary.BigEndian, &buf); err != nil {
		return err
	}

	v.Version = buf[0]
	v.Flags = [3]byte{buf[1], buf[2], buf[3]}

	for _, i := range []interface{}{
		&v.GraphicsMode,
		&v.OpColor[0],
		&v.OpColor[1],
		&v.OpColor[2],
	} {
		if err := binary.Read(r, binary.BigEndian, i); err != nil {
			return err
		}
	}

	return nil
}
