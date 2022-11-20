package iso14496

import (
	"io"
)

type FREE struct{}

func (f *FREE) Parse(r io.ReadSeeker, l int) error {
	if l == 8 {
		return nil
	}

	_, err := r.Seek(int64(l)-8, io.SeekCurrent)
	if err != nil {
		return err
	}

	return nil
}
