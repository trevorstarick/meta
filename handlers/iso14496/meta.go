package iso14496

import (
	"io"
)

type META struct{}

func (m *META) Parse(r io.ReadSeeker, l int) error {
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
		default:
			if err := debug(r, "meta", name, length); err != nil {
				return err
			}
		}
	}

	return nil
}
