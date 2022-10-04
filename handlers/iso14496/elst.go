package iso14496

import (
	"encoding/binary"
	"io"
)

type ELST struct {
	Version byte
	Flags   [3]byte
	Entries int32

	ELT []struct {
		TrackDuration int32
		MediaTime     int32
		MediaRate     int32
	}
}

func (e *ELST) Parse(r io.ReadSeeker, l int) error {
	var buf [4]byte

	if err := binary.Read(r, binary.BigEndian, &buf); err != nil {
		return err
	}

	e.Version = buf[0]
	e.Flags = [3]byte{buf[1], buf[2], buf[3]}

	if err := binary.Read(r, binary.BigEndian, &e.Entries); err != nil {
		return err
	}

	var elt struct {
		TrackDuration int32
		MediaTime     int32
		MediaRate     int32
	}

	for i := 0; i < int(e.Entries); i++ {
		for _, j := range []interface{}{
			&elt.TrackDuration,
			&elt.MediaTime,
			&elt.MediaRate,
		} {
			if err := binary.Read(r, binary.BigEndian, j); err != nil {
				return err
			}
		}

		e.ELT = append(e.ELT, elt)
	}

	return nil
}
