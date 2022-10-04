package iso14496

import (
	"io"
)

type UDTA struct{}

func (u *UDTA) Parse(r io.ReadSeeker, l int) error {
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
		// case "meta":
		// 	meta := &META{}
		// 	err := meta.Parse(r, length)
		// 	if err != nil {
		// 		return err
		// 	}
		default:
			if err := debug(r, "udta", name, length); err != nil {
				return err
			}
		}
	}

	return nil
}
