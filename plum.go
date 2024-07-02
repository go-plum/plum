package plum

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync"
)

type Plum struct {
	Router
	opts serverOptions
	pool sync.Pool
	mux  *http.ServeMux
	srv  *http.Server

	RemoteIPHeaders []string
}

func New(opt ...ServerOption) *Plum {
	opts := defaultServerOptions
	for _, o := range opt {
		o.apply(&opts)
	}

	p := &Plum{
		opts: opts,
		Router: Router{
			basePath: "/",
		},
		RemoteIPHeaders: []string{"X-Forwarded-For", "X-Real-IP"},
		mux:             http.NewServeMux(),
	}
	p.Use(Recover)

	p.pool.New = func() any {
		return p.allocateContext()
	}
	p.Router.engine = p

	RoutePerf(&p.Router)
	return p
}

func (p *Plum) Run(addr string, server ...*http.Server) error {
	p.srv = &http.Server{
		Handler:           p,
		Addr:              addr,
		ReadHeaderTimeout: p.opts.readHeaderTimeout,
	}
	if len(server) != 0 {
		p.srv = server[0]
	}
	return p.srv.ListenAndServe()
}

func (p *Plum) RunTLS(addr, certFile, keyFile string, server ...*http.Server) error {
	p.srv = &http.Server{
		Handler:           p,
		Addr:              addr,
		ReadHeaderTimeout: p.opts.readHeaderTimeout,
	}
	if len(server) != 0 {
		p.srv = server[0]
	}
	return p.srv.ListenAndServeTLS(certFile, keyFile)
}

func (p *Plum) RunServer(lis net.Listener, server *http.Server) error {
	if server == nil {
		return errors.New("plum: no server")
	}
	server.Handler = p
	p.srv = server
	return p.srv.Serve(lis)
}

// Shutdown the http server without interrupting active connections.
func (p *Plum) Shutdown(ctx context.Context) error {
	if p.srv == nil {
		return errors.New("plum: no server")
	}
	return p.srv.Shutdown(ctx)
}

// ServeHTTP should write reply headers and data to the ResponseWriter and then return.
func (p *Plum) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.RequestURI == "*" {
		if req.ProtoAtLeast(1, 1) {
			res.Header().Set("Connection", "close")
		}
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	h, pt := p.mux.Handler(req)
	if pt == "" {
		fmt.Println("not found ") // TODO NOT FOUND
		return
	}
	h.ServeHTTP(res, req)
}

func (p *Plum) allocateContext() *Context {
	return &Context{}
}
