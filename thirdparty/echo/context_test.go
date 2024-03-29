package echo

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"text/template"
	"time"

	testify "github.com/stretchr/testify/assert"
	"github.com/zerjioang/gotools/lib/codes"
	"github.com/zerjioang/gotools/thirdparty/gommon/log"
)

type (
	Template struct {
		templates *template.Template
	}
)

var testUser = user{1, "Jon Snow"}

func BenchmarkAllocJSONP(b *testing.B) {
	e := New()
	req := httptest.NewRequest(POST, "/", strings.NewReader(userJSON))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec).(*HttpContext)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		c.JSONP(codes.StatusOK, "callback", testUser)
	}
}

func BenchmarkAllocJSON(b *testing.B) {
	e := New()
	req := httptest.NewRequest(POST, "/", strings.NewReader(userJSON))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec).(*HttpContext)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		c.JSON(codes.StatusOK, testUser)
	}
}

func BenchmarkAllocXML(b *testing.B) {
	e := New()
	req := httptest.NewRequest(POST, "/", strings.NewReader(userJSON))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec).(*HttpContext)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		c.XML(codes.StatusOK, testUser)
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type responseWriterErr struct {
}

func (responseWriterErr) Header() http.Header {
	return http.Header{}
}

func (responseWriterErr) Write([]byte) (int, error) {
	return 0, errors.New("err")
}

func (responseWriterErr) WriteHeader(statusCode int) {

}

func TestContext(t *testing.T) {
	e := New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec).(*HttpContext)

	assert := testify.New(t)

	// Echo
	assert.Equal(e, c.Echo())

	// Request
	assert.NotNil(c.Request())

	// Response
	assert.NotNil(c.Response())

	//--------
	// Render
	//--------

	tmpl := &Template{
		templates: template.Must(template.New("hello").Parse("Hello, {{.}}!")),
	}
	c.echo.Renderer = tmpl
	err := c.Render(codes.StatusOK, "hello", "Jon Snow")
	if assert.NoError(err) {
		assert.Equal(codes.StatusOK, rec.Code)
		assert.Equal("Hello, Jon Snow!", rec.Body.String())
	}

	c.echo.Renderer = nil
	err = c.Render(codes.StatusOK, "hello", "Jon Snow")
	assert.Error(err)

	// JSON
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec).(*HttpContext)
	err = c.JSON(codes.StatusOK, user{1, "Jon Snow"})
	if assert.NoError(err) {
		assert.Equal(codes.StatusOK, rec.Code)
		assert.Equal(MIMEApplicationJSONCharsetUTF8, rec.Header().Get(HeaderContentType))
		assert.Equal(userJSON+"\n", rec.Body.String())
	}

	// JSON with "?pretty"
	req = httptest.NewRequest(http.MethodGet, "/?pretty", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec).(*HttpContext)
	err = c.JSON(codes.StatusOK, user{1, "Jon Snow"})
	if assert.NoError(err) {
		assert.Equal(codes.StatusOK, rec.Code)
		assert.Equal(MIMEApplicationJSONCharsetUTF8, rec.Header().Get(HeaderContentType))
		assert.Equal(userJSONPretty+"\n", rec.Body.String())
	}
	req = httptest.NewRequest(http.MethodGet, "/", nil) // reset

	// JSONPretty
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec).(*HttpContext)
	err = c.JSONPretty(codes.StatusOK, user{1, "Jon Snow"}, "  ")
	if assert.NoError(err) {
		assert.Equal(codes.StatusOK, rec.Code)
		assert.Equal(MIMEApplicationJSONCharsetUTF8, rec.Header().Get(HeaderContentType))
		assert.Equal(userJSONPretty+"\n", rec.Body.String())
	}

	// JSON (error)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec).(*HttpContext)
	err = c.JSON(codes.StatusOK, make(chan bool))
	assert.Error(err)

	// JSONP
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec).(*HttpContext)
	callback := "callback"
	err = c.JSONP(codes.StatusOK, callback, user{1, "Jon Snow"})
	if assert.NoError(err) {
		assert.Equal(codes.StatusOK, rec.Code)
		assert.Equal(MIMEApplicationJavaScriptCharsetUTF8, rec.Header().Get(HeaderContentType))
		assert.Equal(callback+"("+userJSON+"\n);", rec.Body.String())
	}

	// XML
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec).(*HttpContext)
	err = c.XML(codes.StatusOK, user{1, "Jon Snow"})
	if assert.NoError(err) {
		assert.Equal(codes.StatusOK, rec.Code)
		assert.Equal(MIMEApplicationXMLCharsetUTF8, rec.Header().Get(HeaderContentType))
		assert.Equal(xml.Header+userXML, rec.Body.String())
	}

	// XML with "?pretty"
	req = httptest.NewRequest(http.MethodGet, "/?pretty", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec).(*HttpContext)
	err = c.XML(codes.StatusOK, user{1, "Jon Snow"})
	if assert.NoError(err) {
		assert.Equal(codes.StatusOK, rec.Code)
		assert.Equal(MIMEApplicationXMLCharsetUTF8, rec.Header().Get(HeaderContentType))
		assert.Equal(xml.Header+userXMLPretty, rec.Body.String())
	}
	req = httptest.NewRequest(http.MethodGet, "/", nil)

	// XML (error)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec).(*HttpContext)
	err = c.XML(codes.StatusOK, make(chan bool))
	assert.Error(err)

	// XML response write error
	c = e.NewContext(req, rec).(*HttpContext)
	c.response.Writer = responseWriterErr{}
	err = c.XML(0, 0)
	testify.Error(t, err)

	// XMLPretty
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec).(*HttpContext)
	err = c.XMLPretty(codes.StatusOK, user{1, "Jon Snow"}, "  ")
	if assert.NoError(err) {
		assert.Equal(codes.StatusOK, rec.Code)
		assert.Equal(MIMEApplicationXMLCharsetUTF8, rec.Header().Get(HeaderContentType))
		assert.Equal(xml.Header+userXMLPretty, rec.Body.String())
	}

	t.Run("empty indent", func(t *testing.T) {
		var (
			u           = user{1, "Jon Snow"}
			buf         = new(bytes.Buffer)
			emptyIndent = ""
		)

		t.Run("json", func(t *testing.T) {
			buf.Reset()
			assert := testify.New(t)

			// New JSONBlob with empty indent
			rec = httptest.NewRecorder()
			c = e.NewContext(req, rec).(*HttpContext)
			enc := json.NewEncoder(buf)
			enc.SetIndent(emptyIndent, emptyIndent)
			err = enc.Encode(u)
			err = c.json(codes.StatusOK, user{1, "Jon Snow"}, emptyIndent)
			if assert.NoError(err) {
				assert.Equal(codes.StatusOK, rec.Code)
				assert.Equal(MIMEApplicationJSONCharsetUTF8, rec.Header().Get(HeaderContentType))
				assert.Equal(buf.String(), rec.Body.String())
			}
		})

		t.Run("xml", func(t *testing.T) {
			buf.Reset()
			assert := testify.New(t)

			// New XMLBlob with empty indent
			rec = httptest.NewRecorder()
			c = e.NewContext(req, rec).(*HttpContext)
			enc := xml.NewEncoder(buf)
			enc.Indent(emptyIndent, emptyIndent)
			err = enc.Encode(u)
			err = c.XMLPretty(codes.StatusOK, user{1, "Jon Snow"}, emptyIndent)
			if assert.NoError(err) {
				assert.Equal(codes.StatusOK, rec.Code)
				assert.Equal(MIMEApplicationXMLCharsetUTF8, rec.Header().Get(HeaderContentType))
				assert.Equal(xml.Header+buf.String(), rec.Body.String())
			}
		})
	})

	// Legacy JSONBlob
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec).(*HttpContext)
	data, err := json.Marshal(user{1, "Jon Snow"})
	assert.NoError(err)
	err = c.JSONBlob(codes.StatusOK, data)
	if assert.NoError(err) {
		assert.Equal(codes.StatusOK, rec.Code)
		assert.Equal(MIMEApplicationJSONCharsetUTF8, rec.Header().Get(HeaderContentType))
		assert.Equal(userJSON, rec.Body.String())
	}

	// Legacy JSONPBlob
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec).(*HttpContext)
	callback = "callback"
	data, err = json.Marshal(user{1, "Jon Snow"})
	assert.NoError(err)
	err = c.JSONPBlob(codes.StatusOK, callback, data)
	if assert.NoError(err) {
		assert.Equal(codes.StatusOK, rec.Code)
		assert.Equal(MIMEApplicationJavaScriptCharsetUTF8, rec.Header().Get(HeaderContentType))
		assert.Equal(callback+"("+userJSON+");", rec.Body.String())
	}

	// Legacy XMLBlob
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec).(*HttpContext)
	data, err = xml.Marshal(user{1, "Jon Snow"})
	assert.NoError(err)
	err = c.XMLBlob(codes.StatusOK, data)
	if assert.NoError(err) {
		assert.Equal(codes.StatusOK, rec.Code)
		assert.Equal(MIMEApplicationXMLCharsetUTF8, rec.Header().Get(HeaderContentType))
		assert.Equal(xml.Header+userXML, rec.Body.String())
	}

	// String
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec).(*HttpContext)
	err = c.String(codes.StatusOK, "Hello, World!")
	if assert.NoError(err) {
		assert.Equal(codes.StatusOK, rec.Code)
		assert.Equal(MIMETextPlainCharsetUTF8, rec.Header().Get(HeaderContentType))
		assert.Equal("Hello, World!", rec.Body.String())
	}

	// HTML
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec).(*HttpContext)
	err = c.HTML(codes.StatusOK, "Hello, <strong>World!</strong>")
	if assert.NoError(err) {
		assert.Equal(codes.StatusOK, rec.Code)
		assert.Equal(MIMETextHTMLCharsetUTF8, rec.Header().Get(HeaderContentType))
		assert.Equal("Hello, <strong>World!</strong>", rec.Body.String())
	}

	// Stream
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec).(*HttpContext)
	r := strings.NewReader("response from a stream")
	err = c.Stream(codes.StatusOK, "application/octet-stream", r)
	if assert.NoError(err) {
		assert.Equal(codes.StatusOK, rec.Code)
		assert.Equal("application/octet-stream", rec.Header().Get(HeaderContentType))
		assert.Equal("response from a stream", rec.Body.String())
	}

	// Attachment
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec).(*HttpContext)
	err = c.Attachment("_fixture/images/walle.png", "walle.png")
	if assert.NoError(err) {
		assert.Equal(codes.StatusOK, rec.Code)
		assert.Equal("attachment; filename=\"walle.png\"", rec.Header().Get(HeaderContentDisposition))
		assert.Equal(219885, rec.Body.Len())
	}

	// Inline
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec).(*HttpContext)
	err = c.Inline("_fixture/images/walle.png", "walle.png")
	if assert.NoError(err) {
		assert.Equal(codes.StatusOK, rec.Code)
		assert.Equal("inline; filename=\"walle.png\"", rec.Header().Get(HeaderContentDisposition))
		assert.Equal(219885, rec.Body.Len())
	}

	// NoContent
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec).(*HttpContext)
	c.NoContent(http.StatusOK)
	assert.Equal(http.StatusOK, rec.Code)

	// Error
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec).(*HttpContext)
	c.Error(errors.New("error"))
	assert.Equal(http.StatusInternalServerError, rec.Code)

	// Reset
	c.SetParamNames("foo")
	c.SetParamValues("bar")
	c.Set("foe", "ban")
	c.query = url.Values(map[string][]string{"fon": {"baz"}})
	c.Reset()
	assert.Equal(0, len(c.ParamValues()))
	assert.Equal(0, len(c.ParamNames()))
	assert.Equal(0, len(c.store))
	assert.Equal("", c.Path())
	assert.Equal(0, len(c.QueryParams()))
}

func TestContext_JSON_CommitsCustomResponseCode(t *testing.T) {
	e := New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec).(*HttpContext)
	err := c.JSON(http.StatusCreated, user{1, "Jon Snow"})

	assert := testify.New(t)
	if assert.NoError(err) {
		assert.Equal(http.StatusCreated, rec.Code)
		assert.Equal(MIMEApplicationJSONCharsetUTF8, rec.Header().Get(HeaderContentType))
		assert.Equal(userJSON+"\n", rec.Body.String())
	}
}

func TestContext_JSON_DoesntCommitResponseCodePrematurely(t *testing.T) {
	e := New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec).(*HttpContext)
	err := c.JSON(http.StatusCreated, map[string]float64{"a": math.NaN()})

	assert := testify.New(t)
	if assert.Error(err) {
		assert.False(c.response.Committed)
	}
}

func TestContextCookie(t *testing.T) {
	e := New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	theme := "theme=light"
	user := "user=Jon Snow"
	req.Header.Add(HeaderCookie, theme)
	req.Header.Add(HeaderCookie, user)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec).(*HttpContext)

	assert := testify.New(t)

	// Read single
	cookie, err := c.Cookie("theme")
	if assert.NoError(err) {
		assert.Equal("theme", cookie.Name)
		assert.Equal("light", cookie.Value)
	}

	// Read multiple
	for _, cookie := range c.Cookies() {
		switch cookie.Name {
		case "theme":
			assert.Equal("light", cookie.Value)
		case "user":
			assert.Equal("Jon Snow", cookie.Value)
		}
	}

	// Write
	cookie = &http.Cookie{
		Name:     "SSID",
		Value:    "Ap4PGTEq",
		Domain:   "labstack.com",
		Path:     "/",
		Expires:  time.Now(),
		Secure:   true,
		HttpOnly: true,
	}
	c.SetCookie(cookie)
	assert.Contains(rec.Header().Get(HeaderSetCookie), "SSID")
	assert.Contains(rec.Header().Get(HeaderSetCookie), "Ap4PGTEq")
	assert.Contains(rec.Header().Get(HeaderSetCookie), "labstack.com")
	assert.Contains(rec.Header().Get(HeaderSetCookie), "Secure")
	assert.Contains(rec.Header().Get(HeaderSetCookie), "HttpOnly")
}

func TestContextPath(t *testing.T) {
	e := New()
	r := e.Router()

	r.Add(http.MethodGet, "/users/:id", nil)
	c := e.NewContext(nil, nil)
	r.Find(http.MethodGet, "/users/1", c)

	assert := testify.New(t)

	assert.Equal("/users/:id", c.Path())

	r.Add(http.MethodGet, "/users/:uid/files/:fid", nil)
	c = e.NewContext(nil, nil)
	r.Find(http.MethodGet, "/users/1/files/1", c)
	assert.Equal("/users/:uid/files/:fid", c.Path())
}

func TestContextPathParam(t *testing.T) {
	e := New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	c := e.NewContext(req, nil)

	// ParamNames
	c.SetParamNames("uid", "fid")
	testify.EqualValues(t, []string{"uid", "fid"}, c.ParamNames())

	// ParamValues
	c.SetParamValues("101", "501")
	testify.EqualValues(t, []string{"101", "501"}, c.ParamValues())

	// Param
	testify.Equal(t, "501", c.Param("fid"))
	testify.Equal(t, "", c.Param("undefined"))
}

func TestContextFormValue(t *testing.T) {
	f := make(url.Values)
	f.Set("name", "Jon Snow")
	f.Set("email", "jon@labstack.com")

	e := New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
	req.Header.Add(HeaderContentType, MIMEApplicationForm)
	c := e.NewContext(req, nil)

	// FormValue
	testify.Equal(t, "Jon Snow", c.FormValue("name"))
	testify.Equal(t, "jon@labstack.com", c.FormValue("email"))

	// FormParams
	params, err := c.FormParams()
	if testify.NoError(t, err) {
		testify.Equal(t, url.Values{
			"name":  []string{"Jon Snow"},
			"email": []string{"jon@labstack.com"},
		}, params)
	}

	// Multipart FormParams error
	req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
	req.Header.Add(HeaderContentType, MIMEMultipartForm)
	c = e.NewContext(req, nil)
	params, err = c.FormParams()
	testify.Nil(t, params)
	testify.Error(t, err)
}

func TestContextQueryParam(t *testing.T) {
	q := make(url.Values)
	q.Set("name", "Jon Snow")
	q.Set("email", "jon@labstack.com")
	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
	e := New()
	c := e.NewContext(req, nil)

	// QueryParam
	testify.Equal(t, "Jon Snow", c.QueryParam("name"))
	testify.Equal(t, "jon@labstack.com", c.QueryParam("email"))

	// QueryParams
	testify.Equal(t, url.Values{
		"name":  []string{"Jon Snow"},
		"email": []string{"jon@labstack.com"},
	}, c.QueryParams())
}

func TestContextFormFile(t *testing.T) {
	e := New()
	buf := new(bytes.Buffer)
	mr := multipart.NewWriter(buf)
	w, err := mr.CreateFormFile("file", "test")
	if testify.NoError(t, err) {
		w.Write([]byte("test"))
	}
	mr.Close()
	req := httptest.NewRequest(http.MethodPost, "/", buf)
	req.Header.Set(HeaderContentType, mr.FormDataContentType())
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	f, err := c.FormFile("file")
	if testify.NoError(t, err) {
		testify.Equal(t, "test", f.Filename)
	}
}

func TestContextMultipartForm(t *testing.T) {
	e := New()
	buf := new(bytes.Buffer)
	mw := multipart.NewWriter(buf)
	mw.WriteField("name", "Jon Snow")
	mw.Close()
	req := httptest.NewRequest(http.MethodPost, "/", buf)
	req.Header.Set(HeaderContentType, mw.FormDataContentType())
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	f, err := c.MultipartForm()
	if testify.NoError(t, err) {
		testify.NotNil(t, f)
	}
}

func TestContextRedirect(t *testing.T) {
	e := New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	testify.Equal(t, nil, c.Redirect(http.StatusMovedPermanently, "http://labstack.github.io/echo"))
	testify.Equal(t, http.StatusMovedPermanently, rec.Code)
	testify.Equal(t, "http://labstack.github.io/echo", rec.Header().Get(HeaderLocation))
	testify.Error(t, c.Redirect(310, "http://labstack.github.io/echo"))
}

func TestContextStore(t *testing.T) {
	var c Context
	c = new(HttpContext)
	c.Set("name", "Jon Snow")
	testify.Equal(t, "Jon Snow", c.Get("name"))
}

func BenchmarkContext_Store(b *testing.B) {
	e := &Echo{}

	c := &HttpContext{
		echo: e,
	}

	for n := 0; n < b.N; n++ {
		c.Set("name", "Jon Snow")
		if c.Get("name") != "Jon Snow" {
			b.Fail()
		}
	}
}

func TestContextHandler(t *testing.T) {
	e := New()
	r := e.Router()
	b := new(bytes.Buffer)

	r.Add(http.MethodGet, "/handler", func(Context) error {
		_, err := b.Write([]byte("handler"))
		return err
	})
	c := e.NewContext(nil, nil)
	r.Find(http.MethodGet, "/handler", c)
	err := c.Handler()(c)
	testify.Equal(t, "handler", b.String())
	testify.NoError(t, err)
}

func TestContext_SetHandler(t *testing.T) {
	var c Context
	c = new(HttpContext)

	testify.Nil(t, c.Handler())

	c.SetHandler(func(c Context) error {
		return nil
	})
	testify.NotNil(t, c.Handler())
}

func TestContext_Path(t *testing.T) {
	path := "/pa/th"

	var c Context
	c = new(HttpContext)

	c.SetPath(path)
	testify.Equal(t, path, c.Path())
}

type validator struct{}

func (*validator) Validate(i interface{}) error {
	return nil
}

func TestContext_Validate(t *testing.T) {
	e := New()
	c := e.NewContext(nil, nil)

	testify.Error(t, c.Validate(struct{}{}))

	e.Validator = &validator{}
	testify.NoError(t, c.Validate(struct{}{}))
}

func TestContext_QueryString(t *testing.T) {
	e := New()

	queryString := "query=string&var=val"

	req := httptest.NewRequest(GET, "/?"+queryString, nil)
	c := e.NewContext(req, nil)

	testify.Equal(t, queryString, c.QueryString())
}

func TestContext_Request(t *testing.T) {
	var c Context
	c = new(HttpContext)

	testify.Nil(t, c.Request())

	req := httptest.NewRequest(GET, "/path", nil)
	c.SetRequest(req)

	testify.Equal(t, req, c.Request())
}

func TestContext_Scheme(t *testing.T) {
	tests := []struct {
		c Context
		s string
	}{
		{
			&HttpContext{
				request: &http.Request{
					TLS: &tls.ConnectionState{},
				},
			},
			"https",
		},
		{
			&HttpContext{
				request: &http.Request{
					Header: http.Header{HeaderXForwardedProto: []string{"https"}},
				},
			},
			"https",
		},
		{
			&HttpContext{
				request: &http.Request{
					Header: http.Header{HeaderXForwardedProtocol: []string{"http"}},
				},
			},
			"http",
		},
		{
			&HttpContext{
				request: &http.Request{
					Header: http.Header{HeaderXForwardedSsl: []string{"on"}},
				},
			},
			"https",
		},
		{
			&HttpContext{
				request: &http.Request{
					Header: http.Header{HeaderXUrlScheme: []string{"https"}},
				},
			},
			"https",
		},
		{
			&HttpContext{
				request: &http.Request{},
			},
			"http",
		},
	}

	for _, tt := range tests {
		testify.Equal(t, tt.s, tt.c.Scheme())
	}
}

func TestContext_IsWebSocket(t *testing.T) {
	tests := []struct {
		c  Context
		ws testify.BoolAssertionFunc
	}{
		{
			&HttpContext{
				request: &http.Request{
					Header: http.Header{HeaderUpgrade: []string{"websocket"}},
				},
			},
			testify.True,
		},
		{
			&HttpContext{
				request: &http.Request{
					Header: http.Header{HeaderUpgrade: []string{"Websocket"}},
				},
			},
			testify.True,
		},
		{
			&HttpContext{
				request: &http.Request{},
			},
			testify.False,
		},
		{
			&HttpContext{
				request: &http.Request{
					Header: http.Header{HeaderUpgrade: []string{"other"}},
				},
			},
			testify.False,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			tt.ws(t, tt.c.IsWebSocket())
		})
	}
}

func TestContext_Bind(t *testing.T) {
	e := New()
	req := httptest.NewRequest(POST, "/", strings.NewReader(userJSON))
	c := e.NewContext(req, nil)
	u := new(user)

	req.Header.Add(HeaderContentType, MIMEApplicationJSON)
	err := c.Bind(u)
	testify.NoError(t, err)
	testify.Equal(t, &user{1, "Jon Snow"}, u)
}

func TestContext_Logger(t *testing.T) {
	e := New()
	c := e.NewContext(nil, nil)

	log1 := c.Logger()
	testify.NotNil(t, log1)

	log2 := log.New("echo2")
	c.SetLogger(log2)
	testify.Equal(t, log2, c.Logger())

	// Resetting the context returns the initial logger
	c.Reset()
	testify.Equal(t, log1, c.Logger())
}

func TestContext_RealIP(t *testing.T) {
	tests := []struct {
		c Context
		s string
	}{
		{
			&HttpContext{
				request: &http.Request{
					Header: http.Header{HeaderXForwardedFor: []string{"127.0.0.1, 127.0.1.1, "}},
				},
			},
			"127.0.0.1",
		},
		{
			&HttpContext{
				request: &http.Request{
					Header: http.Header{
						"X-Real-Ip": []string{"192.168.0.1"},
					},
				},
			},
			"192.168.0.1",
		},
		{
			&HttpContext{
				request: &http.Request{
					RemoteAddr: "89.89.89.89:1654",
				},
			},
			"89.89.89.89",
		},
	}

	for _, tt := range tests {
		testify.Equal(t, tt.s, tt.c.RealIP())
	}
}
