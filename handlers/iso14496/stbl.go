package iso14496

import (
	"io"
)

type STBL struct{}

func (s *STBL) Parse(r io.ReadSeeker, l int) error {
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
		case "stsd":
			stsd := &STSD{}
			err := stsd.Parse(r, length)
			if err != nil {
				return err
			}
		// case "stts":
		// 	stts := &STTS{}
		// 	err := stts.Parse(r, length)
		// 	if err != nil {
		// 		return err
		// 	}
		// case "stss":
		// 	stss := &STSS{}
		// 	err := stss.Parse(r, length)
		// 	if err != nil {
		// 		return err
		// 	}
		// case "stsc":
		// 	stsc := &STSC{}
		// 	err := stsc.Parse(r, length)
		// 	if err != nil {
		// 		return err
		// 	}
		// case "stsz":
		// 	stsz := &STSZ{}
		// 	err := stsz.Parse(r, length)
		// 	if err != nil {
		// 		return err
		// 	}
		default:
			if err := debug(r, "stbl", name, length); err != nil {
				return err
			}
		}
	}

	return nil
}
