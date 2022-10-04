package iso14496

import (
	"io"
)

type MDIA struct{}

func (u *MDIA) Parse(r io.ReadSeeker, l int) error {
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
			mdhd := &MDHD{}
			err := mdhd.Parse(r, length)
			if err != nil {
				return err
			}
		case "hdlr":
			hdlr := &HDLR{}
			err := hdlr.Parse(r, length)
			if err != nil {
				return err
			}
		case "minf":
			minf := &MINF{}
			err := minf.Parse(r, length)
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
