package plum

import (
	"net/http"
	"net/http/pprof"
)

const (
	// DefaultPrefix url prefix of pprof
	DefaultPrefix = "/debug/pprof"
)

func getPrefix(prefixOptions ...string) string {
	prefix := DefaultPrefix
	if len(prefixOptions) > 0 {
		prefix = prefixOptions[0]
	}
	return prefix
}

// RoutePerf the standard HandlerFuncs from the net/http/pprof package with
// the provided gin.GrouterGroup. prefixOptions is a optional. If not prefixOptions,
// the default path prefix is used, otherwise first prefixOptions will be path prefix.
func RoutePerf(rg *Router, prefixOptions ...string) {
	prefix := getPrefix(prefixOptions...)

	prefixRouter := rg.Group(prefix)
	{
		prefixRouter.GET("/", pprofHandlerFunc(pprof.Index))
		prefixRouter.GET("/cmdline", pprofHandlerFunc(pprof.Cmdline))
		prefixRouter.GET("/profile", pprofHandlerFunc(pprof.Profile))
		prefixRouter.POST("/symbol", pprofHandlerFunc(pprof.Symbol))
		prefixRouter.GET("/symbol", pprofHandlerFunc(pprof.Symbol))
		prefixRouter.GET("/trace", pprofHandlerFunc(pprof.Trace))

		prefixRouter.GET("/allocs", pprofHandler(pprof.Handler("allocs")))
		prefixRouter.GET("/block", pprofHandler(pprof.Handler("block")))
		prefixRouter.GET("/goroutine", pprofHandler(pprof.Handler("goroutine")))
		prefixRouter.GET("/heap", pprofHandler(pprof.Handler("heap")))
		prefixRouter.GET("/mutex", pprofHandler(pprof.Handler("mutex")))
		prefixRouter.GET("/threadcreate", pprofHandler(pprof.Handler("threadcreate")))
	}
}

func pprofHandlerFunc(h http.HandlerFunc) HandlerFunc {
	return func(ctx *Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func pprofHandler(h http.Handler) HandlerFunc {
	return func(ctx *Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
