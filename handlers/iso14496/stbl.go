package iso14496

import (
	"io"
)

type STBL struct {
	STSD STSD
}

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
			err := s.STSD.Parse(r, length)
			if err != nil {
				return err
			}
		// case "stts":
		// 	err := s.STTS.Parse(r, length)
		// 	if err != nil {
		// 		return err
		// 	}
		// case "stss":
		// 	err := s.STSS.Parse(r, length)
		// 	if err != nil {
		// 		return err
		// 	}
		// case "stsc":
		// 	err := s.STSC.Parse(r, length)
		// 	if err != nil {
		// 		return err
		// 	}
		// case "stsz":
		// 	err := s.STSZ.Parse(r, length)
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
