package iso14496

import (
	"io"
)

type TRAK struct {
	TKHD TKHD
	EDTS EDTS
	MDIA MDIA
}

func (t *TRAK) Parse(r io.ReadSeeker, l int) error {
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
		case "tkhd":
			err := t.TKHD.Parse(r, length)
			if err != nil {
				return err
			}
		case "edts":
			err := t.EDTS.Parse(r, length)
			if err != nil {
				return err
			}
		case "mdia":
			err := t.MDIA.Parse(r, length)
			if err != nil {
				return err
			}
		default:
			if err := debug(r, "trak", name, length); err != nil {
				return err
			}
		}
	}

	return nil
}
