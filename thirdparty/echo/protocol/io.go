package protocol

//serialization function header

const (
	ContentTypeAll      = "*/*"
	ContentTypeHtml     = "text/html"
	ContentTypeJson     = "application/json"
	ContentTypeXml      = "application/xml"
	ContentTypeMsgPack  = "application/x-msgpack"
	ContentTypeProtoBuf = "application/protobuf"
)

type ContentTypeMode uint8

func (m ContentTypeMode) String() string {
	return contentTypeNames[m]
}

var (
	contentTypeNames = []string{
		ContentTypeAll,
		ContentTypeHtml,
		ContentTypeJson,
		ContentTypeXml,
		ContentTypeMsgPack,
		ContentTypeProtoBuf,
	}
)

const (
	ModeAll ContentTypeMode = iota
	ModeHtml
	ModeJson
	ModeXML
	ModeMsgPack
	ModeProtoBuff
)
