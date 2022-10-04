package meta

import (
	"os"

	"github.com/trevorstarick/meta/structs"
)

func Extract(f string) (m *structs.Meta, err error) {
	m = &structs.Meta{}

	r, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	sig, err := GetSignature(r)
	if err != nil {
		return nil, err
	}

	m.Type = sig

	for _, hdlr := range handlerLUT[sig] {
		if err = hdlr(r, m); err != nil {
			return nil, err
		}
	}

	return m, nil
}
