package iso14496

import (
	"io"
)

type EDTS struct{}

func (u *EDTS) Parse(r io.ReadSeeker, l int) error {
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
		case "elst":
			elst := &ELST{}
			err := elst.Parse(r, length)
			if err != nil {
				return err
			}
		default:
			if err := debug(r, "edts", name, length); err != nil {
				return err
			}
		}
	}

	return nil
}
