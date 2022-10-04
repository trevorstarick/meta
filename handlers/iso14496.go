package handlers

import (
	"errors"
	"io"

	"github.com/trevorstarick/meta/handlers/iso14496"
	"github.com/trevorstarick/meta/structs"
)

func parseAtoms(r io.ReadSeeker, m *structs.Meta) error {
	for {
		if err := parseAtom(r, m); err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}

			return err
		}
	}
}

func parseAtom(r io.ReadSeeker, m *structs.Meta) error {
	length, name, err := iso14496.GetAtom(r)
	if err != nil {
		return err
	}

	switch string(name) {
	case "ftyp":
		ftyp := &iso14496.FTYP{}
		err := ftyp.Parse(r, length)
		if err != nil {
			return err
		}
	case "moov":
		moov := &iso14496.MOOV{}
		err := moov.Parse(r, length)
		if err != nil {
			return err
		}
	default:
		if err := iso14496.Debug(r, "root", name, length); err != nil {
			return err
		}
	}

	return nil
}

func ISO14496(r io.ReadSeeker, m *structs.Meta) error {
	defer func() {
		m.Extractors = append(m.Extractors, "ISO14496")
	}()
	// Skip to the beginning of the file.
	if _, err := r.Seek(0, io.SeekStart); err != nil {
		return err
	}

	if err := parseAtoms(r, m); err != nil {
		return err
	}

	return nil
}
