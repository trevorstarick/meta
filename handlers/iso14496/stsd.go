package iso14496

import (
	"encoding/binary"
	"io"
)

type STSD struct {
	Version    byte
	Flags      [3]byte
	EntryCount int32
	Entries    []struct {
		Size               int32
		Format             [4]byte
		Reserved           [6]byte
		DataReferenceIndex int16
		Data               []byte // todo: parse this
	}
}

func (s *STSD) Parse(r io.ReadSeeker, l int) error {
	var buf [4]byte

	if err := binary.Read(r, binary.BigEndian, &buf); err != nil {
		return err
	}

	s.Version = buf[0]
	s.Flags = [3]byte{buf[1], buf[2], buf[3]}

	if err := binary.Read(r, binary.BigEndian, &s.EntryCount); err != nil {
		return err
	}

	for i := 0; i < int(s.EntryCount); i++ {
		entry := struct {
			Size               int32
			Format             [4]byte
			Reserved           [6]byte
			DataReferenceIndex int16
			Data               []byte
		}{}

		if err := binary.Read(r, binary.BigEndian, &entry.Size); err != nil {
			return err
		}

		if err := binary.Read(r, binary.BigEndian, &entry.Format); err != nil {
			return err
		}

		if err := binary.Read(r, binary.BigEndian, &entry.Reserved); err != nil {
			return err
		}

		if err := binary.Read(r, binary.BigEndian, &entry.DataReferenceIndex); err != nil {
			return err
		}

		entry.Data = make([]byte, entry.Size-16)
		if err := binary.Read(r, binary.BigEndian, &entry.Data); err != nil {
			return err
		}

		s.Entries = append(s.Entries, entry)
	}

	return nil
}
