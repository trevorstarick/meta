package iso14496

import (
	"encoding/binary"
	"io"
)

type FTYP struct {
	Brand            string
	MinorVersion     int32
	CompatibleBrands []string
}

func (f *FTYP) Parse(r io.ReadSeeker, l int) error {
	var buf [4]byte

	if err := binary.Read(r, binary.BigEndian, &buf); err != nil {
		return err
	}
	f.Brand = string(buf[:])

	if err := binary.Read(r, binary.BigEndian, &f.MinorVersion); err != nil {
		return err
	}

	l -= 8 + 4 + 4 // length - (atom + brand + minor version)

	for l > 0 {
		if err := binary.Read(r, binary.BigEndian, &buf); err != nil {
			return err
		}

		f.CompatibleBrands = append(f.CompatibleBrands, string(buf[:]))

		l -= 4
	}

	return nil
}
