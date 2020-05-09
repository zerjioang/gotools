package common

// this functions convert input go object to byte array
type Serializer func(v interface{}) ([]byte, error)
