package module

import (
	"io"
)

type CodecType interface {
	NewEncoder(w io.Writer) Encoder
	NewDecoder(r io.Reader) Decoder
}

type Encoder interface {
	Encode(msg interface{}) error
}

type Decoder interface {
	Decode(msg interface{}) error
}


