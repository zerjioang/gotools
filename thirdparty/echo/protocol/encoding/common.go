package encoding

import (
	"strings"

	"github.com/zerjioang/gotools/common"
	"github.com/zerjioang/gotools/thirdparty/echo/protocol"
	"github.com/zerjioang/gotools/thirdparty/echo/protocol/encoding/gogoproto"
	"github.com/zerjioang/gotools/thirdparty/echo/protocol/encoding/html"
	"github.com/zerjioang/gotools/thirdparty/echo/protocol/encoding/json"
	"github.com/zerjioang/gotools/thirdparty/echo/protocol/encoding/msgpack"
	"github.com/zerjioang/gotools/thirdparty/echo/protocol/encoding/xml"
)

var (
	internalErrorHtml      = `<html>error</html>`
	internalErrorHtmlBytes = []byte(internalErrorHtml)
)

// return appropiate content type as descriped in HTTP header value Content-Type
func EncodingSelector(contentType string) (common.Serializer, protocol.ContentTypeMode) {
	switch contentType {
	case protocol.ContentTypeAll, protocol.ContentTypeJson:
		return json.Serialize, protocol.ModeJson
	case protocol.ContentTypeXml:
		return xml.Serialize, protocol.ModeXML
	case protocol.ContentTypeMsgPack:
		return msgpack.Serialize, protocol.ModeMsgPack
	case protocol.ContentTypeProtoBuf:
		return gogoproto.Serialize, protocol.ModeProtoBuff
	case protocol.ContentTypeHtml, hasHtml(contentType):
		return html.Serialize, protocol.ModeHtml
	default:
		//return json serializer as default when no one matches
		return json.Serialize, protocol.ModeJson
	}
}

func hasHtml(contentType string) string {
	if strings.Contains(contentType, protocol.ContentTypeHtml) {
		return contentType
	}
	return ""
}

// return appropiate content type as descriped in HTTP header value Content-Type
func EncodingModeSelector(mode protocol.ContentTypeMode) (common.Serializer, protocol.ContentTypeMode) {
	switch mode {
	case protocol.ModeJson:
		return json.Serialize, protocol.ModeJson
	case protocol.ModeXML:
		return xml.Serialize, protocol.ModeXML
	case protocol.ModeMsgPack:
		return msgpack.Serialize, protocol.ModeMsgPack
	case protocol.ModeProtoBuff:
		return gogoproto.Serialize, protocol.ModeProtoBuff
	default:
		//return json serializer as default when no one matches
		return json.Serialize, protocol.ModeJson
	}
}

// this method uses user requested encoding serializer via http headers
// and encodes result as byte array
func GetContentTypedBytes(contentType string, v interface{}) []byte {
	srlzr, _ := EncodingSelector(contentType)
	return GetBytesFromSerializer(srlzr, v)
}

func GetBytesFromSerializer(s common.Serializer, v interface{}) []byte {
	raw, _ := s(v)
	return raw
}

func GetBytesFromMode(mode protocol.ContentTypeMode, v interface{}) []byte {
	raw, _ := EncodingModeSelector(mode)
	data, _ := raw(v)
	return data
}
