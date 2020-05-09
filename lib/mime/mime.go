package mime

const (
	charsetUTF8 = "charset=UTF-8"
)

type MimePointer uint8

const (
	MimeApplicationJSON MimePointer = iota
	MimeApplicationJSONCharsetUTF8
	MimeApplicationJavaScript
	MimeApplicationJavaScriptCharsetUTF8
	MimeApplicationXML
	MimeApplicationXMLCharsetUTF8
	MimeTextXML
	MimeTextXMLCharsetUTF8
	MimeApplicationForm
	MimeApplicationProtobuf
	MimeApplicationMsgpack
	MimeTextHTML
	MimeTextHTMLCharsetUTF8
	MimeTextPlain
	MimeTextPlainCharsetUTF8
	MimeMultipartForm
	MimeOctetStream
)

// internally used mime types
const (
	mimeApplicationJSON                  = "application/json"
	mimeApplicationJSONCharsetUTF8       = mimeApplicationJSON + "; " + charsetUTF8
	mimeApplicationJavaScript            = "application/javascript"
	mimeApplicationJavaScriptCharsetUTF8 = mimeApplicationJavaScript + "; " + charsetUTF8
	mimeApplicationXML                   = "application/xml"
	mimeApplicationXMLCharsetUTF8        = mimeApplicationXML + "; " + charsetUTF8
	mimeTextXML                          = "text/xml"
	mimeTextXMLCharsetUTF8               = mimeTextXML + "; " + charsetUTF8
	mimeApplicationForm                  = "application/x-www-form-urlencoded"
	mimeApplicationProtobuf              = "application/protobuf"
	mimeApplicationMsgpack               = "application/msgpack"
	mimeTextHTML                         = "text/html"
	mimeTextHTMLCharsetUTF8              = mimeTextHTML + "; " + charsetUTF8
	mimeTextPlain                        = "text/plain"
	mimeTextPlainCharsetUTF8             = mimeTextPlain + "; " + charsetUTF8
	mimeMultipartForm                    = "multipart/form-data"
	mimeOctetStream                      = "application/octet-stream"
)

var (
	MimeList = [17]string{
		mimeApplicationJSON,
		mimeApplicationJSONCharsetUTF8,
		mimeApplicationJavaScript,
		mimeApplicationJavaScriptCharsetUTF8,
		mimeApplicationXML,
		mimeApplicationXMLCharsetUTF8,
		mimeTextXML,
		mimeTextXMLCharsetUTF8,
		mimeApplicationForm,
		mimeApplicationProtobuf,
		mimeApplicationMsgpack,
		mimeTextHTML,
		mimeTextHTMLCharsetUTF8,
		mimeTextPlain,
		mimeTextPlainCharsetUTF8,
		mimeMultipartForm,
		mimeOctetStream,
	}
)

func ToMimetype(p MimePointer) string {
	return MimeList[p]
}
