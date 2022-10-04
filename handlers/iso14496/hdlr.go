package iso14496

import (
	"encoding/binary"
	"io"
)

type HDLR struct {
	Version               byte
	Flags                 [3]byte
	ComponentType         int32
	ComponentSubType      int32
	ComponentManufacturer int32
	ComponentFlags        int32
	ComponentFlagsMask    int32
	ComponentName         string
}

func (h *HDLR) Parse(r io.ReadSeeker, l int) error {
	var buf [4]byte

	if err := binary.Read(r, binary.BigEndian, &buf); err != nil {
		return err
	}

	h.Version = buf[0]
	h.Flags = [3]byte{buf[1], buf[2], buf[3]}

	for _, i := range []interface{}{
		&h.ComponentType,
		&h.ComponentSubType,
		&h.ComponentManufacturer,
		&h.ComponentFlags,
		&h.ComponentFlagsMask,
	} {
		if err := binary.Read(r, binary.BigEndian, i); err != nil {
			return err
		}
	}

	rem := l
	rem -= 8     // atom
	rem -= 4     // version + flags
	rem -= 4 * 5 // component, sub, manufacturer, flags, mask

	vbuf := make([]byte, rem)
	if err := binary.Read(r, binary.BigEndian, &vbuf); err != nil {
		return err
	}

	h.ComponentName = string(vbuf)

	return nil
}
