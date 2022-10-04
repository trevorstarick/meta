package iso14496

import (
	"encoding/binary"
	"io"
)

type MDHD struct {
	Version          byte
	Flags            [3]byte
	CreationTime     int32
	ModificationTime int32
	Timescale        int32
	Duration         int32
	Language         int16
	Quality          int16
}

func (m *MDHD) Parse(r io.ReadSeeker, l int) error {
	var buf [4]byte

	if err := binary.Read(r, binary.BigEndian, &buf); err != nil {
		return err
	}

	m.Version = buf[0]
	m.Flags = [3]byte{buf[1], buf[2], buf[3]}

	for _, i := range []interface{}{
		&m.CreationTime,
		&m.ModificationTime,
		&m.Timescale,
		&m.Duration,
		&m.Language,
		&m.Quality,
	} {
		if err := binary.Read(r, binary.BigEndian, i); err != nil {
			return err
		}
	}

	return nil
}
