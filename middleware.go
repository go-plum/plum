package plum

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime"
)

// Recover will recover from panics.
func Recover(handler HandlerFunc) HandlerFunc {
	return func(ctx *Context) {
		defer func() {
			var rawReq []byte
			if err := recover(); err != nil {
				buf := make([]byte, 64<<10)
				buf = buf[:runtime.Stack(buf, false)]
				if ctx.Request != nil {
					rawReq, _ = httputil.DumpRequest(ctx.Request, false)
				}
				_, _ = fmt.Fprintf(os.Stderr, "Plum call recovery panic: %s\n%v\n%s\n", string(rawReq), err, buf)
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		handler(ctx)
	}
}
