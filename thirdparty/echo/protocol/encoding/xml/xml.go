package xml

import "encoding/xml"

// must implement Serializer function definition at common.go
func Serialize(v interface{}) ([]byte, error) {
	return xml.MarshalIndent(v, "", "   ")
}
