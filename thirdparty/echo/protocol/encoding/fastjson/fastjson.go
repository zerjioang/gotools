package fastjson

import (
	jsoniter "github.com/json-iterator/go"
)

var (
	jsonfast = jsoniter.ConfigFastest
)

// must implement Serializer function definition at common.go
func Serialize(v interface{}) ([]byte, error) {
	return jsonfast.Marshal(v)
}
