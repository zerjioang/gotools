package middleware

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/gotools/lib/codes"
	"github.com/zerjioang/gotools/thirdparty/echo"
)

func TestBodyDump(t *testing.T) {
	e := echo.New()
	hw := "Hello, World!"
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(hw))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := func(c echo.Context) error {
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}
		return c.String(codes.StatusOK, string(body))
	}

	requestBody := ""
	responseBody := ""
	mw := BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		requestBody = string(reqBody)
		responseBody = string(resBody)
	})

	assert := assert.New(t)

	if assert.NoError(mw(h)(c)) {
		assert.Equal(requestBody, hw)
		assert.Equal(responseBody, hw)
		assert.Equal(codes.StatusOK, rec.Code)
		assert.Equal(hw, rec.Body.String())
	}

	// Must set default skipper
	BodyDumpWithConfig(BodyDumpConfig{
		Skipper: nil,
		Handler: func(c echo.Context, reqBody, resBody []byte) {
			requestBody = string(reqBody)
			responseBody = string(resBody)
		},
	})
}

func TestBodyDumpFails(t *testing.T) {
	e := echo.New()
	hw := "Hello, World!"
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(hw))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := func(c echo.Context) error {
		return errors.New("some error")
	}

	mw := BodyDump(func(c echo.Context, reqBody, resBody []byte) {})

	if !assert.Error(t, mw(h)(c)) {
		t.FailNow()
	}

	assert.Panics(t, func() {
		mw = BodyDumpWithConfig(BodyDumpConfig{
			Skipper: nil,
			Handler: nil,
		})
	})

	assert.NotPanics(t, func() {
		mw = BodyDumpWithConfig(BodyDumpConfig{
			Skipper: func(c echo.Context) bool {
				return true
			},
			Handler: func(c echo.Context, reqBody, resBody []byte) {
			},
		})

		if !assert.Error(t, mw(h)(c)) {
			t.FailNow()
		}
	})
}
