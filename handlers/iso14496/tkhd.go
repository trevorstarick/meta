package iso14496

import (
	"encoding/binary"
	"io"
)

type TKHD struct {
	Version          byte
	Flags            [3]byte
	CreationTime     int32
	ModificationTime int32
	TrackID          int32
	_Reserved        int32
	Duration         int32
	_Reserved2       [8]byte
	Layer            int16
	AlternateGroup   int16
	Volume           int16
	_Reserved3       int16
	Matrix           [36]byte
	Width            int32
	Height           int32
}

func (t *TKHD) Parse(r io.ReadSeeker, l int) error {
	var buf [4]byte

	if err := binary.Read(r, binary.BigEndian, &buf); err != nil {
		return err
	}

	t.Version = buf[0]
	t.Flags = [3]byte{buf[1], buf[2], buf[3]}

	for _, i := range []interface{}{
		&t.CreationTime,
		&t.ModificationTime,
		&t.TrackID,
		&t._Reserved,
		&t.Duration,
		&t._Reserved2,
		&t.Layer,
		&t.AlternateGroup,
		&t.Volume,
		&t._Reserved3,
		&t.Matrix,
		&t.Width,
		&t.Height,
	} {
		if err := binary.Read(r, binary.BigEndian, i); err != nil {
			return err
		}
	}

	return nil
}
