package echo

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/zerjioang/gotools/lib/codes"
	"github.com/zerjioang/gotools/thirdparty/echo/protocol"
)

// Context represents the context of the current HTTP request. It holds request and
// response objects, path, path parameters, data and registered handler.
type Context interface {
	// Request returns `*http.Request`.
	Request() *http.Request

	// SetRequest sets `*http.Request`.
	SetRequest(r *http.Request)

	// SetResponse sets `*Response`.
	SetResponse(r *Response)

	// Response returns `*Response`.
	Response() *Response

	// IsTLS returns true if HTTP connection is TLS otherwise false.
	IsTLS() bool

	// IsWebSocket returns true if HTTP connection is WebSocket otherwise false.
	IsWebSocket() bool

	// Scheme returns the HTTP protocol scheme, `http` or `https`.
	Scheme() protocol.RequestScheme

	// RealIP returns the client's network address based on `X-Forwarded-For`
	// or `X-Real-IP` request header.
	RealIP() string

	// Path returns the registered path for the handler.
	Path() string

	// SetPath sets the registered path for the handler.
	SetPath(p string)

	// Param returns path parameter by name.
	Param(name string) string

	// ParamNames returns path parameter names.
	ParamNames() []string

	// SetParamNames sets path parameter names.
	SetParamNames(names ...string)

	// ParamValues returns path parameter values.
	ParamValues() []string

	// SetParamValues sets path parameter values.
	SetParamValues(values ...string)

	// QueryParam returns the query param for the provided name.
	QueryParam(name string) string

	// QueryParams returns the query parameters as `url.Values`.
	QueryParams() url.Values

	// QueryString returns the URL query string.
	QueryString() string

	// FormValue returns the form field value for the provided name.
	FormValue(name string) string

	// FormParams returns the form parameters as `url.Values`.
	FormParams() (url.Values, error)

	// FormFile returns the multipart form file for the provided name.
	FormFile(name string) (*multipart.FileHeader, error)

	// MultipartForm returns the multipart form.
	MultipartForm() (*multipart.Form, error)

	// Cookie returns the named cookie provided in the request.
	Cookie(name string) (*http.Cookie, error)

	// SetCookie adds a `Set-Cookie` header in HTTP response.
	SetCookie(cookie *http.Cookie)

	// Cookies returns the HTTP cookies sent with the request.
	Cookies() []*http.Cookie

	// Get retrieves data from the context.
	Get(key string) interface{}

	// Set saves data in the context.
	Set(key string, val interface{})

	// Bind binds the request body into provided type `i`. The default binder
	// does it based on Content-Type header.
	Bind(i interface{}) error

	// Validate validates provided `i`. It is usually called after `Context#Bind()`.
	// Validator must be registered using `Echo#Validator`.
	Validate(i interface{}) error

	// Render renders a template with data and sends a text/html response with status
	// code. Renderer must be registered using `Echo.Renderer`.
	Render(code codes.HttpStatusCode, name string, data interface{}) error

	// HTML sends an HTTP response with status code.
	HTML(code codes.HttpStatusCode, html string) error

	// HTMLBlob sends an HTTP blob response with status code.
	HTMLBlob(code codes.HttpStatusCode, b []byte) error

	// String sends a string response with status code.
	String(code codes.HttpStatusCode, s string) error

	// JSON sends a JSON response with status code.
	JSON(code codes.HttpStatusCode, i interface{}) error

	// JSONPretty sends a pretty-print JSON with status code.
	JSONPretty(code codes.HttpStatusCode, i interface{}, indent string) error

	// JSONBlob sends a JSON blob response with status code.
	JSONBlob(code codes.HttpStatusCode, b []byte) error

	// JSONP sends a JSONP response with status code. It uses `callback` to construct
	// the JSONP payload.
	JSONP(code codes.HttpStatusCode, callback string, i interface{}) error

	// JSONPBlob sends a JSONP blob response with status code. It uses `callback`
	// to construct the JSONP payload.
	JSONPBlob(code codes.HttpStatusCode, callback string, b []byte) error

	// XML sends an XML response with status code.
	XML(code codes.HttpStatusCode, i interface{}) error

	// XMLPretty sends a pretty-print XML with status code.
	XMLPretty(code codes.HttpStatusCode, i interface{}, indent string) error

	// XMLBlob sends an XML blob response with status code.
	XMLBlob(code codes.HttpStatusCode, b []byte) error

	// Blob sends a blob response with status code and content type.
	Blob(code codes.HttpStatusCode, contentType string, b []byte) error

	// Stream sends a streaming response with status code and content type.
	Stream(code codes.HttpStatusCode, contentType string, r io.Reader) error

	// File sends a response with the content of the file.
	File(file string) error

	// Attachment sends a response as attachment, prompting client to save the
	// file.
	Attachment(file string, name string) error

	// Inline sends a response as inline, opening the file in the browser.
	Inline(file string, name string) error

	// NoContent sends a response with no body and a status code.
	NoContent(code codes.HttpStatusCode) error

	// Redirect redirects the request to a provided URL with status code.
	Redirect(code codes.HttpStatusCode, url string) error

	// Error invokes the registered HTTP error handler. Generally used by middleware.
	Error(err error)

	// Handler returns the matched handler by router.
	Handler() HandlerFunc

	// SetHandler sets the matched handler by router.
	SetHandler(h HandlerFunc)

	// Logger returns the `Logger` instance.
	Logger() Logger

	// Set the logger
	SetLogger(l Logger)

	// Echo returns the `Echo` instance.
	Echo() *Echo

	// Reset resets the context after request completes. It must be called along
	// with `Echo#AcquireContext()` and `Echo#ReleaseContext()`.
	// See `Echo#ServeHTTP()`
	//Reset(r *http.Request, w http.ResponseWriter)
	Reset()

	Preload(r *http.Request, w http.ResponseWriter)
	WriteContentType(value string)
	Body() []byte
	RequestContentType() string
}
