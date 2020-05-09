package str

import "encoding/json"

func GetJsonBytes(o interface{}) {
	json.Marshal(o)
}

func StdMarshal(o interface{}) ([]byte, error) {
	return json.Marshal(o)
}
