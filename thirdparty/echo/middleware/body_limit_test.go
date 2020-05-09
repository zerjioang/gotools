package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zerjioang/gotools/lib/parsers/bytesize"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/gotools/lib/codes"
	"github.com/zerjioang/gotools/thirdparty/echo"
)

func TestBodyLimit(t *testing.T) {
	e := echo.New()
	hw := []byte("Hello, World!")
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(hw))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := func(c echo.Context) error {
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}
		return c.String(codes.StatusOK, string(body))
	}

	assert := assert.New(t)

	// Based on content length (within limit)
	if assert.NoError(BodyLimit(bytesize.FromHumanSizeSilent("2M"))(h)(c)) {
		assert.Equal(codes.StatusOK, rec.Code)
		assert.Equal(hw, rec.Body.Bytes())
	}

	// Based on content read (overlimit)
	he := BodyLimit(bytesize.FromHumanSizeSilent("2B"))(h)(c).(*echo.HTTPError)
	assert.Equal(codes.StatusRequestEntityTooLarge, he.Code)

	// Based on content read (within limit)
	req = httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(hw))
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	if assert.NoError(BodyLimit(bytesize.FromHumanSizeSilent("2M"))(h)(c)) {
		assert.Equal(codes.StatusOK, rec.Code)
		assert.Equal("Hello, World!", rec.Body.String())
	}

	// Based on content read (overlimit)
	req = httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(hw))
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	he = BodyLimit(bytesize.FromHumanSizeSilent("2B"))(h)(c).(*echo.HTTPError)
	assert.Equal(codes.StatusRequestEntityTooLarge, he.Code)
}

func TestBodyLimitReader(t *testing.T) {
	hw := []byte("Hello, World!")
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(hw))
	rec := httptest.NewRecorder()

	config := BodyLimitConfig{
		Skipper: DefaultSkipper,
		limit:   2,
	}
	reader := &limitedReader{
		BodyLimitConfig: config,
		reader:          ioutil.NopCloser(bytes.NewReader(hw)),
		context:         e.NewContext(req, rec),
	}

	// read all should return ErrStatusRequestEntityTooLarge
	_, err := ioutil.ReadAll(reader)
	he := err.(*echo.HTTPError)
	assert.Equal(t, codes.StatusRequestEntityTooLarge, he.Code)

	// reset reader and read two bytes must succeed
	bt := make([]byte, 2)
	reader.Reset(ioutil.NopCloser(bytes.NewReader(hw)), e.NewContext(req, rec))
	n, err := reader.Read(bt)
	assert.Equal(t, 2, n)
	assert.Equal(t, nil, err)
}
