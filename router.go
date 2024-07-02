package plum

import (
	"fmt"
	"net/http"
	"slices"
	"strings"
)

type Router struct {
	scope       string
	basePath    string
	engine      *Plum
	middlewares []Middleware
}

func (r *Router) Group(relativePath string, m ...Middleware) *Router {
	newScope := r.scope + relativePath
	newRouter := &Router{
		scope:       newScope,
		basePath:    joinPaths(r.basePath, relativePath),
		engine:      r.engine,
		middlewares: r.middlewares,
	}
	if len(m) > 0 {
		slices.Reverse(m)
		newRouter.middlewares = slices.Concat(m, newRouter.middlewares)
	}
	return newRouter
}
func (r *Router) Use(m ...Middleware) {
	slices.Reverse(m)
	r.middlewares = slices.Concat(m, r.middlewares)
}

func (r *Router) withMiddlewares(handler HandlerFunc) HandlerFunc {
	for _, middleware := range r.middlewares {
		handler = middleware(handler)
	}
	return handler
}

func (r *Router) POST(route string, handler HandlerFunc) {
	r.Handle(http.MethodPost, route, handler)
}

func (r *Router) GET(route string, handler HandlerFunc) {
	r.Handle(http.MethodGet, route, handler)
}

func (r *Router) Handle(method, route string, handler HandlerFunc) {
	if strings.HasSuffix(route, "/") {
		route += "{$}"
	}
	rh := &RouterHandler{
		engine: r.engine,
		h:      r.withMiddlewares(handler),
	}
	fmt.Println(method + " " + r.scope + route)
	r.engine.mux.Handle(method+" "+r.scope+route, rh)
}
