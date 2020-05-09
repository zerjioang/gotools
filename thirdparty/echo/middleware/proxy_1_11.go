// +build go1.11

package middleware

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/zerjioang/gotools/lib/codes"
	"github.com/zerjioang/gotools/thirdparty/echo"
)

func proxyHTTP(tgt *ProxyTarget, c echo.Context, config ProxyConfig) http.Handler {
	proxy := httputil.NewSingleHostReverseProxy(tgt.URL)
	proxy.ErrorHandler = func(resp http.ResponseWriter, req *http.Request, err error) {
		desc := tgt.URL.String()
		if tgt.Name != "" {
			desc = fmt.Sprintf("%s(%s)", tgt.Name, tgt.URL.String())
		}
		c.Set("_error", echo.NewHTTPError(codes.StatusBadGateway, fmt.Sprintf("remote %s unreachable, could not forward: %v", desc, err)))
	}
	proxy.Transport = config.Transport
	return proxy
}
