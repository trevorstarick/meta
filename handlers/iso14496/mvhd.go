package iso14496

import (
	"encoding/binary"
	"io"
)

type MVHD struct {
	Version           byte
	Flags             [3]byte
	CreationTime      int32
	ModificationTime  int32
	TimeScale         int32
	Duration          int32
	PrefRate          int32
	PrefVolume        int16
	Reserved          [10]byte
	Matrix            [36]byte
	PreviewTime       int32
	PreviewDuration   int32
	PosterTime        int32
	SelectionTime     int32
	SelectionDuration int32
	CurrentTime       int32
	NextTrackID       int32
}

func (m *MVHD) Parse(r io.ReadSeeker, l int) error {
	var buf [4]byte

	if err := binary.Read(r, binary.BigEndian, &buf); err != nil {
		return err
	}

	m.Version = buf[0]
	m.Flags = [3]byte{buf[1], buf[2], buf[3]}

	for _, i := range []interface{}{
		&m.CreationTime,
		&m.ModificationTime,
		&m.TimeScale,
		&m.Duration,
		&m.PrefRate,
		&m.PrefVolume,
		&m.Reserved,
		&m.Matrix,
		&m.PreviewTime,
		&m.PreviewDuration,
		&m.PosterTime,
		&m.SelectionTime,
		&m.SelectionDuration,
		&m.CurrentTime,
		&m.NextTrackID,
	} {
		if err := binary.Read(r, binary.BigEndian, i); err != nil {
			return err
		}
	}

	return nil
}
