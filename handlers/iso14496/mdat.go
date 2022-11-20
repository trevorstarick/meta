package iso14496

import (
	"io"
)

type MDAT struct{}

func (f *MDAT) Parse(r io.ReadSeeker, l int) error {
	_, err := r.Seek(int64(l)-8, io.SeekCurrent)
	if err != nil {
		return err
	}

	return nil
}
