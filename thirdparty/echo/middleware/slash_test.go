// Copyright gotools
// SPDX-License-Identifier: GNU GPL v3

package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/gotools/lib/codes"
	"github.com/zerjioang/gotools/thirdparty/echo"
)

func TestAddTrailingSlash(t *testing.T) {
	is := assert.New(t)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/add-slash", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := AddTrailingSlash()(func(c echo.Context) error {
		return nil
	})
	is.NoError(h(c))
	is.Equal("/add-slash/", req.URL.Path)
	is.Equal("/add-slash/", req.RequestURI)

	// Method Connect must not fail:
	req = httptest.NewRequest(http.MethodConnect, "", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	h = AddTrailingSlash()(func(c echo.Context) error {
		return nil
	})
	is.NoError(h(c))
	is.Equal("/", req.URL.Path)
	is.Equal("/", req.RequestURI)

	// With config
	req = httptest.NewRequest(http.MethodGet, "/add-slash?key=value", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	h = AddTrailingSlashWithConfig(TrailingSlashConfig{
		RedirectCode: codes.StatusMovedPermanently.Int(),
	})(func(c echo.Context) error {
		return nil
	})
	is.NoError(h(c))
	is.Equal(codes.StatusMovedPermanently, rec.Code)
	is.Equal("/add-slash/?key=value", rec.Header().Get(echo.HeaderLocation))
}

func TestRemoveTrailingSlash(t *testing.T) {
	is := assert.New(t)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/remove-slash/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := RemoveTrailingSlash()(func(c echo.Context) error {
		return nil
	})
	is.NoError(h(c))
	is.Equal("/remove-slash", req.URL.Path)
	is.Equal("/remove-slash", req.RequestURI)

	// Method Connect must not fail:
	req = httptest.NewRequest(http.MethodConnect, "", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	h = RemoveTrailingSlash()(func(c echo.Context) error {
		return nil
	})
	is.NoError(h(c))
	is.Equal("", req.URL.Path)
	is.Equal("", req.RequestURI)

	// With config
	req = httptest.NewRequest(http.MethodGet, "/remove-slash/?key=value", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	h = RemoveTrailingSlashWithConfig(TrailingSlashConfig{
		RedirectCode: codes.StatusMovedPermanently,
	})(func(c echo.Context) error {
		return nil
	})
	is.NoError(h(c))
	is.Equal(codes.StatusMovedPermanently, rec.Code)
	is.Equal("/remove-slash?key=value", rec.Header().Get(echo.HeaderLocation))

	// With bare URL
	req = httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	h = RemoveTrailingSlash()(func(c echo.Context) error {
		return nil
	})
	is.NoError(h(c))
	is.Equal("", req.URL.Path)
}
