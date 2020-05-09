// Copyright gotools
// SPDX-License-Identifier: GNU GPL v3

package echo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"

	jsoniter "github.com/json-iterator/go"
	"github.com/zerjioang/gotools/lib/codes"
	"github.com/zerjioang/gotools/lib/concurrentmap"
	"github.com/zerjioang/gotools/lib/logger"
)

var (
	errInvalidateCache = errors.New("failed to get item from internal cache, cache invalidation issues around")
)

// Context represents the Context of the current HTTP request. It holds request and
// response objects, path, path parameters, data and registered handler.
type HttpContext struct {
	Context
	request  *http.Request
	response *Response
	path     string
	pnames   []string
	pvalues  []string
	query    url.Values
	handler  HandlerFunc
	store    Map
	echo     *Echo
	lock     sync.RWMutex
	// client ip
	ip          string
	contentType string
	preloaded   atomic.Value
}

var (
	fileCache concurrentmap.ConcurrentMap
)

func init() {
	fileCache = concurrentmap.New()
}

const (
	defaultMemory = 32 << 20 // 32 MB
	indexPage     = "index.html"
)

func (c *HttpContext) Preload(r *http.Request, w http.ResponseWriter) {
	// re initialize current http request fields
	if !c.preloaded.Load().(bool) {
		c.request = r
		c.response.reset(w)
		//TODO content-type negotiation
		c.contentType = c.request.Header.Get(HeaderAccept)
		c.ip = c.resolveRealIP()
		c.preloaded.Store(true)
	}
}

func (c *HttpContext) writeContentType(value string) {
	header := c.response.Header()
	if header.Get(HeaderContentType) == "" {
		header.Set(HeaderContentType, value)
	}
}

func (c *HttpContext) WriteContentType(value string) {
	c.writeContentType(value)
}

func (c *HttpContext) Request() *http.Request {
	return c.request
}

func (c *HttpContext) SetRequest(r *http.Request) {
	c.request = r
}

func (c *HttpContext) Response() *Response {
	return c.response
}

func (c *HttpContext) resolveRealIP() string {
	if ipstr := c.request.Header.Get(HeaderXForwardedFor); ipstr != "" {
		return strings.Split(ipstr, ", ")[0]
	}
	if ipstr := c.request.Header.Get(HeaderXRealIP); ipstr != "" {
		return ipstr
	}
	ra, _ := c.SplitHostPort(c.request.RemoteAddr)
	return ra
}

// simplistic method that split host and port from string
// the reason for custom method is to avoid
// the overhead of net package and its CGO methods
func (c *HttpContext) SplitHostPort(address string) (string, string) {
	var ipStr, port string
	if address != "" {
		loc := strings.LastIndex(address, ":")
		if loc != -1 {
			ipStr = address[0:loc]
			port = address[loc+1:]
		} else {
			ipStr = address
		}
	}
	return ipStr, port
}

func (c *HttpContext) Path() string {
	return c.path
}

func (c *HttpContext) RealIP() string {
	return c.ip
}

func (c *HttpContext) SetPath(p string) {
	c.path = p
}

func (c *HttpContext) Param(name string) string {
	for i, n := range c.pnames {
		if i < len(c.pvalues) {
			if n == name {
				return c.pvalues[i]
			}
		}
	}
	return ""
}

func (c *HttpContext) ParamNames() []string {
	return c.pnames
}

func (c *HttpContext) SetParamNames(names ...string) {
	c.pnames = names
}

func (c *HttpContext) ParamValues() []string {
	return c.pvalues[:len(c.pnames)]
}

func (c *HttpContext) SetParamValues(values ...string) {
	c.pvalues = values
}

func (c *HttpContext) QueryParam(name string) string {
	if c.query == nil {
		c.query = c.request.URL.Query()
	}
	return c.query.Get(name)
}

func (c *HttpContext) QueryParams() url.Values {
	if c.query == nil {
		c.query = c.request.URL.Query()
	}
	return c.query
}

func (c *HttpContext) QueryString() string {
	return c.request.URL.RawQuery
}

func (c *HttpContext) FormValue(name string) string {
	return c.request.FormValue(name)
}

func (c *HttpContext) FormParams() (url.Values, error) {
	if strings.HasPrefix(c.request.Header.Get(HeaderContentType), MIMEMultipartForm) {
		if err := c.request.ParseMultipartForm(defaultMemory); err != nil {
			return nil, err
		}
	} else {
		if err := c.request.ParseForm(); err != nil {
			return nil, err
		}
	}
	return c.request.Form, nil
}

func (c *HttpContext) FormFile(name string) (*multipart.FileHeader, error) {
	_, fh, err := c.request.FormFile(name)
	return fh, err
}

func (c *HttpContext) MultipartForm() (*multipart.Form, error) {
	err := c.request.ParseMultipartForm(defaultMemory)
	return c.request.MultipartForm, err
}

func (c *HttpContext) Cookie(name string) (*http.Cookie, error) {
	return c.request.Cookie(name)
}

func (c *HttpContext) SetCookie(cookie *http.Cookie) {
	http.SetCookie(c.response, cookie)
}

func (c *HttpContext) Cookies() []*http.Cookie {
	return c.request.Cookies()
}

func (c *HttpContext) Get(key string) interface{} {
	c.lock.RLock()
	v := c.store[key]
	c.lock.RUnlock()
	return v
}

func (c *HttpContext) Set(key string, val interface{}) {
	c.lock.Lock()

	if c.store == nil {
		c.store = make(Map)
	}
	c.store[key] = val
	c.lock.Unlock()
}

func (c *HttpContext) Bind(i interface{}) error {
	return json.NewDecoder(c.request.Body).Decode(i)
}

func (c *HttpContext) Render(code codes.HttpStatusCode, name string, data interface{}) (err error) {
	if c.echo.Renderer == nil {
		return ErrRendererNotRegistered
	}
	buf := new(bytes.Buffer)
	if err = c.echo.Renderer.Render(buf, name, data, c); err != nil {
		return
	}
	return c.HTMLBlob(code, buf.Bytes())
}

func (c *HttpContext) HTML(code codes.HttpStatusCode, html string) (err error) {
	return c.HTMLBlob(code, []byte(html))
}

func (c *HttpContext) HTMLBlob(code codes.HttpStatusCode, b []byte) (err error) {
	return c.Blob(code, MIMETextHTMLCharsetUTF8, b)
}

func (c *HttpContext) String(code codes.HttpStatusCode, s string) (err error) {
	return c.Blob(code, MIMETextPlainCharsetUTF8, []byte(s))
}

func (c *HttpContext) jsonPBlob(code codes.HttpStatusCode, callback string, i interface{}) (err error) {
	enc := json.NewEncoder(c.response)
	_, pretty := c.QueryParams()["pretty"]
	if c.echo.Debug || pretty {
		enc.SetIndent("", "  ")
	}
	c.writeContentType(MIMEApplicationJavaScriptCharsetUTF8)
	c.response.WriteHeaderCode(code)
	if _, err = c.response.Write([]byte(callback + "(")); err != nil {
		return
	}
	if err = enc.Encode(i); err != nil {
		return
	}
	if _, err = c.response.Write([]byte(");")); err != nil {
		return
	}
	return
}

func (c *HttpContext) json(code codes.HttpStatusCode, i interface{}, indent string) error {
	enc := json.NewEncoder(c.response)
	if indent != "" {
		enc.SetIndent("", indent)
	}
	c.writeContentType(MIMEApplicationJSONCharsetUTF8)
	c.response.WriteHeaderCode(code)
	return enc.Encode(i)
}

//custom json encoder
func (c *HttpContext) JSON(code codes.HttpStatusCode, i interface{}) (err error) {
	raw, encErr := jsoniter.Marshal(i)
	if encErr != nil {
		return encErr
	}
	return c.Blob(code, MIMEApplicationJSONCharsetUTF8, raw)
}

func (c *HttpContext) JSONBlob(code codes.HttpStatusCode, b []byte) (err error) {
	return c.Blob(code, MIMEApplicationJSONCharsetUTF8, b)
}

func (c *HttpContext) JSONP(code codes.HttpStatusCode, callback string, i interface{}) (err error) {
	return c.jsonPBlob(code, callback, i)
}

func (c *HttpContext) JSONPBlob(code codes.HttpStatusCode, callback string, b []byte) (err error) {
	c.writeContentType(MIMEApplicationJavaScriptCharsetUTF8)
	c.response.WriteHeaderCode(code)
	if _, err = c.response.Write([]byte(callback + "(")); err != nil {
		return
	}
	if _, err = c.response.Write(b); err != nil {
		return
	}
	_, err = c.response.Write([]byte(");"))
	return
}

func (c *HttpContext) Blob(code codes.HttpStatusCode, contentType string, b []byte) (err error) {
	c.writeContentType(contentType)
	c.response.WriteHeaderCode(code)
	_, err = c.response.Write(b)
	return
}

func (c *HttpContext) Stream(code codes.HttpStatusCode, contentType string, r io.Reader) (err error) {
	c.writeContentType(contentType)
	c.response.WriteHeaderCode(code)
	_, err = io.Copy(c.response, r)
	return
}

func (c *HttpContext) File(file string) (err error) {
	initialFilePath := file
	// check if file is cached
	// check if file was already readed before and saved in our cache
	// this avoid overhead on disk readings
	object, found := fileCache.Get(initialFilePath)
	if found && object != nil {
		// cast
		buffer, ok := object.(*FileBuffer)
		if ok {
			//casting was ok
			// add a http cache directive too
			c.response.Header().Set("Cache-Control", "public, max-age=86400") // 24h cache = 86400
			http.ServeContent(c.response, c.request, buffer.name, buffer.time, buffer)
			return nil
		} else {
			// some cache and data error occured.
			return errInvalidateCache
		}
	} else {
		// file not cached
		f, err := os.Open(file)
		if err != nil {
			return NotFoundHandler(c)
		}

		fi, _ := f.Stat()
		if fi.IsDir() {
			//append index.html if directory detected
			file = filepath.Join(file, indexPage)
		}
		f, err = os.Open(file)
		if err != nil {
			return NotFoundHandler(c)
		}
		if fi, err = f.Stat(); err != nil {
			return f.Close()
		}
		// before sending file data to the client, create a filebuffer  for caching purposes
		raw, _ := ioutil.ReadAll(f)
		b := bytes.Buffer{}
		_, _ = b.Write(raw)
		item := new(FileBuffer)
		item.name = fi.Name()
		item.time = fi.ModTime()
		item.Buffer = b
		item.Index = 0
		fileCache.Set(initialFilePath, item)
		// add a http cache directive too
		c.response.Header().Set("Cache-Control", "public, max-age=86400") // 24h cache = 86400
		http.ServeContent(c.response, c.request, fi.Name(), fi.ModTime(), f)
		return f.Close()
	}
}

func (c *HttpContext) Attachment(file, name string) error {
	return c.contentDisposition(file, name, "attachment")
}

func (c *HttpContext) Inline(file, name string) error {
	return c.contentDisposition(file, name, "inline")
}

func (c *HttpContext) contentDisposition(file, name, dispositionType string) error {
	c.response.Header().Set(HeaderContentDisposition, fmt.Sprintf("%s; filename=%q", dispositionType, name))
	return c.File(file)
}

func (c *HttpContext) NoContent(code codes.HttpStatusCode) error {
	c.response.WriteHeaderCode(code)
	return nil
}

func (c *HttpContext) Redirect(code codes.HttpStatusCode, url string) error {
	if code < 300 || code > 308 {
		return ErrInvalidRedirectCode
	}
	c.response.Header().Set(HeaderLocation, url)
	c.response.WriteHeaderCode(code)
	return nil
}

func (c *HttpContext) Error(err error) {
	c.echo.HTTPErrorHandler(err, c)
}

func (c *HttpContext) Echo() *Echo {
	return c.echo
}

func (c *HttpContext) Handler() HandlerFunc {
	return c.handler
}

func (c *HttpContext) SetHandler(h HandlerFunc) {
	c.handler = h
}

func (c *HttpContext) Logger() Logger {
	return c.echo.Logger
}

func (c *HttpContext) Reset() {
	c.request = nil
	c.response.reset(nil)
	c.query = nil
	c.handler = NotFoundHandler
	c.store = nil
	c.path = ""
	c.pnames = nil
	c.preloaded.Store(false)
	c.ip = ""
}

// the the content of the request body
func (c *HttpContext) Body() []byte {
	var content []byte
	body := c.request.Body
	hasBody := body != nil
	if hasBody {
		var err error
		content, err = ioutil.ReadAll(body)
		if err == nil {
			return content
		}
		logger.Error("failed to read request content body due to: ", err)
	}
	return content
}

func (c *HttpContext) RequestContentType() string {
	return c.contentType
}
