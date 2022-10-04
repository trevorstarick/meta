package meta

import (
	"io"

	"github.com/trevorstarick/meta/errors"
	"github.com/trevorstarick/meta/handlers"
)

type entry struct {
	typ      string
	sig      []byte
	handlers []handlers.Handler
}

type node []node

var (
	tree []node

	maxSigLen  = 0
	sigLUT     = map[string]string{}
	handlerLUT = map[string][]handlers.Handler{}

	filetypes = []entry{
		{
			"mp4-iso",
			[]byte{0x00, 0x00, 0x00, 0x20, 0x66, 0x74, 0x79, 0x70, 0x69, 0x73, 0x6f, 0x6D},
			[]handlers.Handler{
				handlers.ISO14496,
			},
		}, {
			"mp4-mp42",
			[]byte{0x00, 0x00, 0x00, 0x20, 0x66, 0x74, 0x79, 0x70, 0x6D, 0x70, 0x34, 0x32},
			[]handlers.Handler{
				handlers.ISO14496,
			},
		},
	}
)

func buildLUT() {
	for _, ft := range filetypes {
		ent := tree

		for i, byt := range ft.sig {
			if ent[byt] == nil {
				ent[byt] = make([]node, 0xff)
			}

			ent = ent[byt]

			if i > maxSigLen {
				maxSigLen = i + 1
			}
		}

		sigLUT[string(ft.sig)] = ft.typ
		handlerLUT[ft.typ] = ft.handlers
	}
}

func GetSignature(r io.Reader) (string, error) {
	var sig string

	buf := make([]byte, maxSigLen)

	e := tree

	sig = ""

	_, err := r.Read(buf)
	if err != nil {
		return "", err
	}

	for _, b := range buf {
		if e[b] != nil {
			e = e[b]
			sig += string(b)
		} else {
			return "", errors.ErrUnknownType
		}
	}

	if typ, ok := sigLUT[sig]; ok {
		return typ, nil
	}

	return "", errors.ErrUnknownType
}
