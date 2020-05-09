package html

import (
	"errors"
)

var (
	errInvalidHtml = errors.New("invalid html content received for serialization")
)

func Serialize(v interface{}) ([]byte, error) {
	raw, ok := v.(string)
	if ok {
		return []byte(raw), nil
	}
	return nil, errInvalidHtml
}
