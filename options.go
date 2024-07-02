package plum

import (
	"log/slog"
	"os"
	"time"

	"github.com/go-plum/plum/render"
)

type serverOptions struct {
	Log Logger

	MaxMultipartMemory int64
	readHeaderTimeout  time.Duration

	HTMLRender render.HTMLRender
}

const defaultMultipartMemory = 32 << 20 // 32 MB

var defaultServerOptions = serverOptions{
	Log:                slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})),
	MaxMultipartMemory: defaultMultipartMemory,
	readHeaderTimeout:  time.Second * 45,
}

// A ServerOption sets options such as credentials, codec and keepalive parameters, etc.
type ServerOption interface {
	apply(*serverOptions)
}

// funcServerOption wraps a function that modifies serverOptions into an
// implementation of the ServerOption interface.
type funcServerOption struct {
	f func(*serverOptions)
}

func (fdo *funcServerOption) apply(do *serverOptions) {
	fdo.f(do)
}

func newFuncServerOption(f func(*serverOptions)) *funcServerOption {
	return &funcServerOption{
		f: f,
	}
}

// ReadHeaderTimeout  this is newFuncServerOption example.
func ReadHeaderTimeout(d time.Duration) ServerOption {
	return newFuncServerOption(func(o *serverOptions) {
		o.readHeaderTimeout = d
	})
}

// HTMLRender  this is newFuncServerOption example.
func HTMLRender(d render.HTMLRender) ServerOption {
	return newFuncServerOption(func(o *serverOptions) {
		o.HTMLRender = d
	})
}

// WithLogger setting logger .
func WithLogger(log Logger) ServerOption {
	return newFuncServerOption(func(o *serverOptions) {
		o.Log = log
	})
}
