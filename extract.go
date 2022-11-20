package meta

import (
	"os"

	"github.com/pkg/errors"
	"github.com/trevorstarick/meta/structs"
)

func Extract(f string) (m *structs.Meta, err error) {
	m = &structs.Meta{}

	r, err := os.Open(f)
	if err != nil {
		return nil, errors.Wrap(err, "os.Open")
	}
	defer r.Close()

	sig, err := GetSignature(r)
	if err != nil {
		return nil, errors.Wrap(err, "GetSignature")
	}

	m.Type = sig

	for _, hdlr := range handlerLUT[sig] {
		if err = hdlr(r, m); err != nil {
			return nil, errors.Wrap(err, "hdlr")
		}
	}

	return m, nil
}
