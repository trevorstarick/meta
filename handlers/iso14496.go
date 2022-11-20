package handlers

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
	"github.com/trevorstarick/meta/handlers/iso14496"
	"github.com/trevorstarick/meta/structs"
)

func ISO14496(r io.ReadSeeker, m *structs.Meta) error {
	defer func() {
		m.Extractors = append(m.Extractors, "ISO14496")
	}()

	// Skip to the beginning of the file.
	if _, err := r.Seek(0, io.SeekStart); err != nil {
		return errors.Wrap(err, "iso14496: r.Seek")
	}

	i := iso14496.ISO14496{}
	if err := i.Parse(r); err != nil {
		if !errors.Is(err, io.EOF) {
			return errors.Wrap(err, "iso14496: Parse")
		}
	}

	// spew.Dump(i)
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	_ = b
	// fmt.Println(string(b))

	return nil
}
