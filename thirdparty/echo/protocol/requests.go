package protocol

type RequestScheme uint8

func (s RequestScheme) String() string {
	return ""
}

const (
	Http RequestScheme = iota
	Https
	Unix
	Websocket
	Other
)
