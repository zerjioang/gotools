package multihash

import (
	"errors"

	"github.com/zerjioang/gotools/multihash/md2"
)

// definition of supported hashing modes
type HashType uint8

const (
	Unknown HashType = iota
	Md2              = 1
)

// definition of encoder function signature
type encoderFunc func([]byte) ([]byte, error)

var (
	errUnknEnc = errors.New("unknown encoder provided")
	encoders   = [2]encoderFunc{
		unknownEnc,
		md2.Encoder,
	}
)

func Encode(typeof HashType, data []byte) ([]byte, error) {
	m := encoders[typeof]
	return m(data)
}

func Decode() {

}

func Parse() {

}

func unknownEnc([]byte) ([]byte, error) {
	return nil, errUnknEnc
}
