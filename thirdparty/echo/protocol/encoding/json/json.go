package json

import (
	"encoding/json"
)

// must implement Serializer function definition at common.go
func Serialize(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
