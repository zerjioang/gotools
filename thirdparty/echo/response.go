// Copyright gotools
// SPDX-License-Identifier: GNU GPL v3

package echo

import (
	"bufio"
	"net"
	"net/http"

	"github.com/zerjioang/gotools/lib/codes"
)

type (
	// Response wraps an http.ResponseWriter and implements its interface to be used
	// by an HTTP handler to construct an HTTP response.
	// See: https://golang.org/pkg/net/http/#ResponseWriter
	Response struct {
		echo        *Echo
		beforeFuncs []func()
		afterFuncs  []func()
		Writer      http.ResponseWriter
		Status      codes.HttpStatusCode
		Size        int64
		Committed   bool
	}
)

// NewResponse creates a new instance of Response.
func NewResponse(w http.ResponseWriter, e *Echo) (r *Response) {
	return &Response{Writer: w, echo: e}
}

// Header returns the header map for the writer that will be sent by
// WriteHeader. Changing the header after a call to WriteHeader (or Write) has
// no effect unless the modified headers were declared as trailers by setting
// the "Trailer" header before the call to WriteHeader (see example)
// To suppress implicit response headers, set their value to nil.
// Example: https://golang.org/pkg/net/http/#example_ResponseWriter_trailers
func (r *Response) Header() http.Header {
	return r.Writer.Header()
}

func (r *Response) HeaderPtr() *http.Header {
	h := r.Writer.Header()
	return &h
}

// Before registers a function which is called just before the response is written.
func (r *Response) Before(fn func()) {
	r.beforeFuncs = append(r.beforeFuncs, fn)
}

// After registers a function which is called just after the response is written.
// If the `Content-Length` is unknown, none of the after function is executed.
func (r *Response) After(fn func()) {
	r.afterFuncs = append(r.afterFuncs, fn)
}

// WriteHeader sends an HTTP response header with status code. If WriteHeader is
// not called explicitly, the first call to Write will trigger an implicit
// WriteHeader(codes.StatusOK). Thus explicit calls to WriteHeader are mainly
// used to send error codes.
func (r *Response) WriteHeaderCode(code codes.HttpStatusCode) {
	if r.Committed {
		//r.echo.Logger.Warn("response already committed")
		return
	}
	for _, fn := range r.beforeFuncs {
		fn()
	}
	r.Status = code
	r.Writer.WriteHeader(code.Int())
	r.Committed = true
}

// implements interface
func (r *Response) WriteHeader(c int) {
	r.WriteHeaderCode(codes.HttpStatusCode(c))
}

// Write writes the data to the connection as part of an HTTP reply.
func (r *Response) Write(b []byte) (n int, err error) {
	if !r.Committed {
		if r.Status == 0 {
			r.Status = http.StatusOK
		}
		r.WriteHeaderCode(r.Status)
	}
	n, err = r.Writer.Write(b)
	r.Size += int64(n)
	for _, fn := range r.afterFuncs {
		fn()
	}
	return
}

// Flush implements the http.Flusher interface to allow an HTTP handler to flush
// buffered data to the client.
// See [http.Flusher](https://golang.org/pkg/net/http/#Flusher)
func (r *Response) Flush() {
	r.Writer.(http.Flusher).Flush()
}

// Hijack implements the http.Hijacker interface to allow an HTTP handler to
// take over the connection.
// See [http.Hijacker](https://golang.org/pkg/net/http/#Hijacker)
func (r *Response) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return r.Writer.(http.Hijacker).Hijack()
}

func (r *Response) reset(w http.ResponseWriter) {
	r.beforeFuncs = nil
	r.afterFuncs = nil
	r.Writer = w
	r.Size = 0
	r.Status = codes.StatusOK
	r.Committed = false
}
