package iso14496

import (
	"io"
)

type MINF struct {
	VMHD VMHD
	DINF DINF
	STBL STBL
}

func (m *MINF) Parse(r io.ReadSeeker, l int) error {
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
			err := m.VMHD.Parse(r, length)
			if err != nil {
				return err
			}
		case "dinf":
			err := m.DINF.Parse(r, length)
			if err != nil {
				return err
			}
		case "stbl":
			err := m.STBL.Parse(r, length)
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
