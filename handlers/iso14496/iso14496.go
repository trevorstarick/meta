package iso14496

import (
	"io"

	"github.com/pkg/errors"
)

type ISO14496 struct {
	FTYP FTYP
	MOOV MOOV
	FREE FREE
	MDAT MDAT
}

func (i *ISO14496) Parse(r io.ReadSeeker) error {
	for {
		length, name, err := GetAtom(r)
		if err != nil {
			if err == io.EOF {
				return nil
			}

			return errors.Wrap(err, "GetAtom")
		}

		switch string(name) {
		case "ftyp":
			err := i.FTYP.Parse(r, length)
			if err != nil {
				return errors.Wrap(err, "FTYP")
			}
		case "moov":
			err := i.MOOV.Parse(r, length)
			if err != nil {
				return errors.Wrap(err, "MOOV")
			}
		case "free":
			err := i.FREE.Parse(r, length)
			if err != nil {
				return errors.Wrap(err, "FREE")
			}
		case "mdat":
			err := i.MDAT.Parse(r, length)
			if err != nil {
				return errors.Wrap(err, "MDAT")
			}
		default:
			if err := debug(r, "root", name, length); err != nil {
				return errors.Wrap(err, "debug")
			}
		}
	}
}
