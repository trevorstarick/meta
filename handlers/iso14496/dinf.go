package iso14496

import (
	"encoding/binary"
	"io"
)

type DINF struct {
	DataReference struct {
		Size       int32
		Type       [4]byte
		Version    byte
		Flags      [3]byte
		EntryCount int32
		Entries    []struct {
			Size    int32
			Type    [4]byte
			Version byte
			Flags   [3]byte
			Data    string
		}
	}
}

func (d *DINF) Parse(r io.ReadSeeker, l int) error {
	for _, i := range []interface{}{
		&d.DataReference.Size,
		&d.DataReference.Type,
		&d.DataReference.Version,
		&d.DataReference.Flags,
		&d.DataReference.EntryCount,
	} {
		if err := binary.Read(r, binary.BigEndian, i); err != nil {
			return err
		}
	}

	for i := 0; i < int(d.DataReference.EntryCount); i++ {
		var entry struct {
			Size    int32
			Type    [4]byte
			Version byte
			Flags   [3]byte
			Data    string
		}

		for _, i := range []interface{}{
			&entry.Size,
			&entry.Type,
			&entry.Version,
			&entry.Flags,
		} {
			if err := binary.Read(r, binary.BigEndian, i); err != nil {
				return err
			}
		}

		vbuf := make([]byte, entry.Size-12)
		if _, err := r.Read(vbuf); err != nil {
			return err
		}

		entry.Data = string(vbuf)

		d.DataReference.Entries = append(d.DataReference.Entries, entry)
	}

	return nil
}
