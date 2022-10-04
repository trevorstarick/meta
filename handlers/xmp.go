package handlers

import (
	"io"

	"github.com/trevorstarick/meta/structs"
)

func XMP(r io.ReadSeeker, m *structs.Meta) error {
	defer func() {
		m.Extractors = append(m.Extractors, "XMP")
	}()

	return nil
}
