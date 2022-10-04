package iso14496

import (
	"io"
)

type MINF struct{}

func (u *MINF) Parse(r io.ReadSeeker, l int) error {
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
		case "vmhd":
			vmhd := &VMHD{}
			err := vmhd.Parse(r, length)
			if err != nil {
				return err
			}
		case "dinf":
			dinf := &DINF{}
			err := dinf.Parse(r, length)
			if err != nil {
				return err
			}
		case "stbl":
			stbl := &STBL{}
			err := stbl.Parse(r, length)
			if err != nil {
				return err
			}
		default:
			if err := debug(r, "minf", name, length); err != nil {
				return err
			}
		}
	}

	return nil
}
