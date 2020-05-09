// Copyright gotools
// SPDX-License-Identifier: GNU GPL v3

package middleware

import (
	"github.com/zerjioang/gotools/lib/codes"
	"github.com/zerjioang/gotools/thirdparty/echo"
)

type (
	// TrailingSlashConfig defines the config for TrailingSlash middleware.
	TrailingSlashConfig struct {
		// Skipper defines a function to skip middleware.
		Skipper Skipper

		// Status code to be used when redirecting the request.
		// Optional, but when provided the request is redirected using this code.
		RedirectCode codes.HttpStatusCode `yaml:"redirect_code"`
	}
)

var (
	// DefaultTrailingSlashConfig is the default TrailingSlash middleware config.
	DefaultTrailingSlashConfig = TrailingSlashConfig{
		Skipper: DefaultSkipper,
	}
)

// AddTrailingSlash returns a root level (before router) middleware which adds a
// trailing slash to the request `URL#Path`.
//
// Usage `Echo#Pre(AddTrailingSlash())`
func AddTrailingSlash() echo.MiddlewareFunc {
	return AddTrailingSlashWithConfig(DefaultTrailingSlashConfig)
}

// AddTrailingSlashWithConfig returns a AddTrailingSlash middleware with config.
// See `AddTrailingSlash()`.
func AddTrailingSlashWithConfig(config TrailingSlashConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = DefaultTrailingSlashConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			url := req.URL
			path := url.Path
			qs := c.QueryString()
			if !HasSuffix(path, "/") {
				path += "/"
				uri := path
				if qs != "" {
					uri += "?" + qs
				}

				// Redirect
				if config.RedirectCode != 0 {
					return c.Redirect(config.RedirectCode, uri)
				}

				// Forward
				req.RequestURI = uri
				url.Path = path
			}
			return next(c)
		}
	}
}

// HasSuffix tests whether the string s ends with suffix.
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// RemoveTrailingSlash returns a root level (before router) middleware which removes
// a trailing slash from the request URI.
//
// Usage `Echo#Pre(RemoveTrailingSlash())`
func RemoveTrailingSlash() echo.MiddlewareFunc {
	return RemoveTrailingSlashWithConfig(TrailingSlashConfig{})
}

// RemoveTrailingSlashWithConfig returns a RemoveTrailingSlash middleware with config.
// See `RemoveTrailingSlash()`.
func RemoveTrailingSlashWithConfig(config TrailingSlashConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = DefaultTrailingSlashConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			url := req.URL
			path := url.Path
			qs := c.QueryString()
			l := len(path) - 1
			if l > 0 && HasSuffix(path, "/") {
				path = path[:l]
				uri := path
				if qs != "" {
					uri += "?" + qs
				}

				// Redirect
				if config.RedirectCode != 0 {
					return c.Redirect(config.RedirectCode, uri)
				}

				// Forward
				req.RequestURI = uri
				url.Path = path
			}
			return next(c)
		}
	}
}
