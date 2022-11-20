package iso14496

import (
	"io"
)

type MDIA struct {
	MDHD MDHD
	HDLR HDLR
	MINF MINF
}

func (m *MDIA) Parse(r io.ReadSeeker, l int) error {
	for {
		length, name, err := GetAtom(r)
		if err != nil {
			return err
		}

		l -= length
		if l <= 0 {
			_, _ = r.Seek(-8, io.SeekCurrent)

			break
		}

		switch string(name) {
		case "mdhd":
			err := m.MDHD.Parse(r, length)
			if err != nil {
				return err
			}
		case "hdlr":
			err := m.HDLR.Parse(r, length)
			if err != nil {
				return err
			}
		case "minf":
			err := m.MINF.Parse(r, length)
			if err != nil {
				return err
			}
		default:
			if err := debug(r, "mdia", name, length); err != nil {
				return err
			}
		}
	}

	return nil
}
