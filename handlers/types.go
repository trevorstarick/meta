package handlers

import (
	"io"

	"github.com/trevorstarick/meta/structs"
)

type Handler func(io.ReadSeeker, *structs.Meta) error
