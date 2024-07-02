package plum

import (
	"net/http"
)

type HandlerFunc func(*Context)

type Middleware func(HandlerFunc) HandlerFunc

type RouterHandler struct {
	h      HandlerFunc
	engine *Plum
}

func (r *RouterHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := r.engine.pool.Get().(*Context)
	ctx.Writer = w
	ctx.Request = req
	ctx.engine = r.engine
	ctx.reset()

	r.h(ctx)

	r.engine.pool.Put(ctx)
}
