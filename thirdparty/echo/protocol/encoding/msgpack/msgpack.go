package msgpack

import (
	"github.com/vmihailenco/msgpack"
)

// must implement Serializer function definition at common.go
func Serialize(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}
