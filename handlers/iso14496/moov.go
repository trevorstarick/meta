package iso14496

import (
	"io"
)

type MOOV struct{}

func (m *MOOV) Parse(r io.ReadSeeker, l int) error {
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
		case "mvhd":
			mvhd := &MVHD{}
			err := mvhd.Parse(r, length)
			if err != nil {
				return err
			}
		case "trak":
			trak := &TRAK{}
			err := trak.Parse(r, length)
			if err != nil {
				return err
			}
		case "udta":
			udta := &UDTA{}
			err := udta.Parse(r, length)
			if err != nil {
				return err
			}
		default:
			if err := debug(r, "moov", name, length); err != nil {
				return err
			}
		}
	}

	return nil
}
