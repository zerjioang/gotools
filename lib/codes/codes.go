// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package codes

// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
type HttpStatusCode uint16

func (c HttpStatusCode) Int() int {
	return int(c)
}

func (c HttpStatusCode) Text() string {
	return StatusTextOptimized(c)
}

const (
	StatusContinue           HttpStatusCode = 100 // RFC 7231, 6.2.1
	StatusSwitchingProtocols HttpStatusCode = 101 // RFC 7231, 6.2.2
	StatusProcessing         HttpStatusCode = 102 // RFC 2518, 10.1
	StatusEarlyHints         HttpStatusCode = 103

	StatusOK                   HttpStatusCode = 200 // RFC 7231, 6.3.1
	StatusCreated              HttpStatusCode = 201 // RFC 7231, 6.3.2
	StatusAccepted             HttpStatusCode = 202 // RFC 7231, 6.3.3
	StatusNonAuthoritativeInfo HttpStatusCode = 203 // RFC 7231, 6.3.4
	StatusNoContent            HttpStatusCode = 204 // RFC 7231, 6.3.5
	StatusResetContent         HttpStatusCode = 205 // RFC 7231, 6.3.6
	StatusPartialContent       HttpStatusCode = 206 // RFC 7233, 4.1
	StatusMultiStatus          HttpStatusCode = 207 // RFC 4918, 11.1
	StatusAlreadyReported      HttpStatusCode = 208 // RFC 5842, 7.1
	StatusIMUsed               HttpStatusCode = 226 // RFC 3229, 10.4.1

	StatusMultipleChoices   HttpStatusCode = 300 // RFC 7231, 6.4.1
	StatusMovedPermanently  HttpStatusCode = 301 // RFC 7231, 6.4.2
	StatusFound             HttpStatusCode = 302 // RFC 7231, 6.4.3
	StatusSeeOther          HttpStatusCode = 303 // RFC 7231, 6.4.4
	StatusNotModified       HttpStatusCode = 304 // RFC 7232, 4.1
	StatusUseProxy          HttpStatusCode = 305 // RFC 7231, 6.4.5
	_                       HttpStatusCode = 306 // RFC 7231, 6.4.6 (Unused)
	StatusTemporaryRedirect HttpStatusCode = 307 // RFC 7231, 6.4.7
	StatusPermanentRedirect HttpStatusCode = 308 // RFC 7538, 3

	StatusBadRequest                   HttpStatusCode = 400 // RFC 7231, 6.5.1
	StatusUnauthorized                 HttpStatusCode = 401 // RFC 7235, 3.1
	StatusPaymentRequired              HttpStatusCode = 402 // RFC 7231, 6.5.2
	StatusForbidden                    HttpStatusCode = 403 // RFC 7231, 6.5.3
	StatusNotFound                     HttpStatusCode = 404 // RFC 7231, 6.5.4
	StatusMethodNotAllowed             HttpStatusCode = 405 // RFC 7231, 6.5.5
	StatusNotAcceptable                HttpStatusCode = 406 // RFC 7231, 6.5.6
	StatusProxyAuthRequired            HttpStatusCode = 407 // RFC 7235, 3.2
	StatusRequestTimeout               HttpStatusCode = 408 // RFC 7231, 6.5.7
	StatusConflict                     HttpStatusCode = 409 // RFC 7231, 6.5.8
	StatusGone                         HttpStatusCode = 410 // RFC 7231, 6.5.9
	StatusLengthRequired               HttpStatusCode = 411 // RFC 7231, 6.5.10
	StatusPreconditionFailed           HttpStatusCode = 412 // RFC 7232, 4.2
	StatusRequestEntityTooLarge        HttpStatusCode = 413 // RFC 7231, 6.5.11
	StatusRequestURITooLong            HttpStatusCode = 414 // RFC 7231, 6.5.12
	StatusUnsupportedMediaType         HttpStatusCode = 415 // RFC 7231, 6.5.13
	StatusRequestedRangeNotSatisfiable HttpStatusCode = 416 // RFC 7233, 4.4
	StatusExpectationFailed            HttpStatusCode = 417 // RFC 7231, 6.5.14
	StatusTeapot                       HttpStatusCode = 418 // RFC 7168, 2.3.3
	StatusMisdirectedRequest           HttpStatusCode = 421 // RFC 7540, 9.1.2
	StatusUnprocessableEntity          HttpStatusCode = 422 // RFC 4918, 11.2
	StatusLocked                       HttpStatusCode = 423 // RFC 4918, 11.3
	StatusFailedDependency             HttpStatusCode = 424 // RFC 4918, 11.4
	StatusTooEarly                     HttpStatusCode = 425 // RFC 8470, 5.2.
	StatusUpgradeRequired              HttpStatusCode = 426 // RFC 7231, 6.5.15
	StatusPreconditionRequired         HttpStatusCode = 428 // RFC 6585, 3
	StatusTooManyRequests              HttpStatusCode = 429 // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge  HttpStatusCode = 431 // RFC 6585, 5
	StatusUnavailableForLegalReasons   HttpStatusCode = 451 // RFC 7725, 3

	StatusInternalServerError           HttpStatusCode = 500 // RFC 7231, 6.6.1
	StatusNotImplemented                HttpStatusCode = 501 // RFC 7231, 6.6.2
	StatusBadGateway                    HttpStatusCode = 502 // RFC 7231, 6.6.3
	StatusServiceUnavailable            HttpStatusCode = 503 // RFC 7231, 6.6.4
	StatusGatewayTimeout                HttpStatusCode = 504 // RFC 7231, 6.6.5
	StatusHTTPVersionNotSupported       HttpStatusCode = 505 // RFC 7231, 6.6.6
	StatusVariantAlsoNegotiates         HttpStatusCode = 506 // RFC 2295, 8.1
	StatusInsufficientStorage           HttpStatusCode = 507 // RFC 4918, 11.5
	StatusLoopDetected                  HttpStatusCode = 508 // RFC 5842, 7.2
	StatusNotExtended                   HttpStatusCode = 510 // RFC 2774, 7
	StatusNetworkAuthenticationRequired HttpStatusCode = 511 // RFC 6585, 6
)

var statusText = map[HttpStatusCode]string{
	StatusContinue:           "Continue",
	StatusSwitchingProtocols: "Switching Protocols",
	StatusProcessing:         "Processing",

	StatusOK:                   "OK",
	StatusCreated:              "Created",
	StatusAccepted:             "Accepted",
	StatusNonAuthoritativeInfo: "Non-Authoritative Information",
	StatusNoContent:            "No Content",
	StatusResetContent:         "Reset Content",
	StatusPartialContent:       "Partial Content",
	StatusMultiStatus:          "Multi-Status",
	StatusAlreadyReported:      "Already Reported",
	StatusIMUsed:               "IM Used",

	StatusMultipleChoices:   "Multiple Choices",
	StatusMovedPermanently:  "Moved Permanently",
	StatusFound:             "Found",
	StatusSeeOther:          "See Other",
	StatusNotModified:       "Not Modified",
	StatusUseProxy:          "Use Proxy",
	StatusTemporaryRedirect: "Temporary Redirect",
	StatusPermanentRedirect: "Permanent Redirect",

	StatusBadRequest:                   "Bad Request",
	StatusUnauthorized:                 "Unauthorized",
	StatusPaymentRequired:              "Payment Required",
	StatusForbidden:                    "Forbidden",
	StatusNotFound:                     "Not Found",
	StatusMethodNotAllowed:             "Method Not Allowed",
	StatusNotAcceptable:                "Not Acceptable",
	StatusProxyAuthRequired:            "Proxy Authentication Required",
	StatusRequestTimeout:               "Request Timeout",
	StatusConflict:                     "Conflict",
	StatusGone:                         "Gone",
	StatusLengthRequired:               "Length Required",
	StatusPreconditionFailed:           "Precondition Failed",
	StatusRequestEntityTooLarge:        "Request Entity Too Large",
	StatusRequestURITooLong:            "Request URI Too Long",
	StatusUnsupportedMediaType:         "Unsupported Media Type",
	StatusRequestedRangeNotSatisfiable: "Requested Range Not Satisfiable",
	StatusExpectationFailed:            "Expectation Failed",
	StatusTeapot:                       "I'm a teapot",
	StatusMisdirectedRequest:           "Misdirected Request",
	StatusUnprocessableEntity:          "Unprocessable Entity",
	StatusLocked:                       "Locked",
	StatusFailedDependency:             "Failed Dependency",
	StatusTooEarly:                     "Too Early",
	StatusUpgradeRequired:              "Upgrade Required",
	StatusPreconditionRequired:         "Precondition Required",
	StatusTooManyRequests:              "Too Many Requests",
	StatusRequestHeaderFieldsTooLarge:  "Request Header Fields Too Large",
	StatusUnavailableForLegalReasons:   "Unavailable For Legal Reasons",

	StatusInternalServerError:           "Internal Server Error",
	StatusNotImplemented:                "Not Implemented",
	StatusBadGateway:                    "Bad Gateway",
	StatusServiceUnavailable:            "Service Unavailable",
	StatusGatewayTimeout:                "Gateway Timeout",
	StatusHTTPVersionNotSupported:       "HTTP Version Not Supported",
	StatusVariantAlsoNegotiates:         "Variant Also Negotiates",
	StatusInsufficientStorage:           "Insufficient Storage",
	StatusLoopDetected:                  "Loop Detected",
	StatusNotExtended:                   "Not Extended",
	StatusNetworkAuthenticationRequired: "Network Authentication Required",
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code HttpStatusCode) string {
	return statusText[code]
}

// helper methods
func IsInformational(code HttpStatusCode) bool {
	// returns \c true if the given \p code is an informational code
	return code >= 100 && code < 200
}
func IsSuccessful(code HttpStatusCode) bool {
	// returns \c true if the given \p code is a successful code
	return code >= 200 && code < 300
}
func IsRedirection(code HttpStatusCode) bool {
	// returns \c true if the given \p code is a redirectional code
	return code >= 300 && code < 400
}
func IsClientError(code HttpStatusCode) bool {
	// returns \c true if the given \p code is a client error code
	return code >= 400 && code < 500
}
func IsServerError(code HttpStatusCode) bool {
	// returns \c true if the given \p code is a server error code
	return code >= 500 && code < 600
}
func IsError(code HttpStatusCode) bool {
	// returns \c true if the given \p code is any type of error code
	return code >= 400
}

/* Returns the standard HTTP reason phrase for a HTTP status code.
 * param code An HTTP status code.
 * return The standard HTTP reason phrase for the given \p code or \c NULL if no standard
 * phrase for the given \p code is known.
 */
/* Returns the standard HTTP reason phrase for a HTTP status code.
 * param code An HTTP status code.
 * return The standard HTTP reason phrase for the given \p code or \c NULL if no standard
 * phrase for the given \p code is known.

name                           old time/op    new time/op     delta
StatusText/status-text-4         15.5ns ±19%      3.7ns ± 7%   -75.96%  (p=0.000 n=10+9)

name                           old speed      new speed       delta
StatusText/status-text-4       65.1MB/s ±17%  269.6MB/s ± 8%  +314.00%  (p=0.000 n=10+9)
*/
func StatusTextOptimized(code HttpStatusCode) string {
	switch code {
	/*####### 1xx - Informational #######*/
	case StatusContinue:
		return "Continue"
	case StatusSwitchingProtocols:
		return "Switching Protocols"
	case StatusProcessing:
		return "Processing"
	case StatusEarlyHints:
		return "Early Hints"
	/*####### 2xx - Successful #######*/
	case StatusOK:
		return "OK"
	case StatusCreated:
		return "Created"
	case StatusAccepted:
		return "Accepted"
	case StatusNonAuthoritativeInfo:
		return "Non-Authoritative Information"
	case StatusNoContent:
		return "No Content"
	case StatusResetContent:
		return "Reset Content"
	case StatusPartialContent:
		return "Partial Content"
	case StatusMultiStatus:
		return "Multi-Status"
	case StatusAlreadyReported:
		return "Already Reported"
	case StatusIMUsed:
		return "IM Used"

	/*####### 3xx - Redirection #######*/
	case StatusMultipleChoices:
		return "Multiple Choices"
	case StatusMovedPermanently:
		return "Moved Permanently"
	case StatusFound:
		return "Found"
	case StatusSeeOther:
		return "See Other"
	case StatusNotModified:
		return "Not Modified"
	case StatusUseProxy:
		return "Use Proxy"
	case StatusTemporaryRedirect:
		return "Temporary Redirect"
	case StatusPermanentRedirect:
		return "Permanent Redirect"

	/*####### 4xx - client Error #######*/
	case StatusBadRequest:
		return "Bad Request"
	case StatusUnauthorized:
		return "Unauthorized"
	case StatusPaymentRequired:
		return "Payment Required"
	case StatusForbidden:
		return "Forbidden"
	case StatusNotFound:
		return "Not Found"
	case StatusMethodNotAllowed:
		return "Method Not Allowed"
	case StatusNotAcceptable:
		return "Not Acceptable"
	case StatusProxyAuthRequired:
		return "Proxy Authentication Required"
	case StatusRequestTimeout:
		return "Request Timeout"
	case StatusConflict:
		return "Conflict"
	case StatusGone:
		return "Gone"
	case StatusLengthRequired:
		return "Length Required"
	case StatusPreconditionFailed:
		return "Precondition Failed"
	case StatusRequestEntityTooLarge:
		return "Payload Too Large"
	case StatusRequestURITooLong:
		return "URI Too Long"
	case StatusUnsupportedMediaType:
		return "Unsupported Media Type"
	case StatusRequestedRangeNotSatisfiable:
		return "Range Not Satisfiable"
	case StatusExpectationFailed:
		return "Expectation Failed"
	case StatusTeapot:
		return "I'm a teapot"
	case StatusMisdirectedRequest:
		return "Misdirected Request"
	case StatusUnprocessableEntity:
		return "Unprocessable Entity"
	case StatusLocked:
		return "Locked"
	case StatusFailedDependency:
		return "Failed Dependency"
	case StatusTooEarly:
		return "Too Early"
	case StatusUpgradeRequired:
		return "Upgrade Required"
	case StatusPreconditionRequired:
		return "Precondition Required"
	case StatusTooManyRequests:
		return "Too Many Requests"
	case StatusRequestHeaderFieldsTooLarge:
		return "Request Header Fields Too Large"
	case StatusUnavailableForLegalReasons:
		return "Unavailable For Legal Reasons"

	/*####### 5xx - Server Error #######*/
	case StatusInternalServerError:
		return "Internal Server Error"
	case StatusNotImplemented:
		return "Not Implemented"
	case StatusBadGateway:
		return "Bad Gateway"
	case StatusServiceUnavailable:
		return "Service Unavailable"
	case StatusGatewayTimeout:
		return "Gateway Time-out"
	case StatusHTTPVersionNotSupported:
		return "HTTP Version Not Supported"
	case StatusVariantAlsoNegotiates:
		return "Variant Also Negotiates"
	case StatusInsufficientStorage:
		return "Insufficient Storage"
	case StatusLoopDetected:
		return "Loop Detected"
	case StatusNotExtended:
		return "Not Extended"
	case StatusNetworkAuthenticationRequired:
		return "Network Authentication Required"
	default:
		return ""
	}
}
