// Package nrhttptreemux is an adaptation of nrhttprouter
// https://github.com/newrelic/go-agent/blob/master/_integrations/nrhttprouter/nrhttprouter.go
package nrhttptreemux

import (
	"net/http"

	"github.com/dimfeld/httptreemux/v5"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// Router should be used in place of httptreemux.Mux and httptreemux.ContextMux.
type Router struct {
	*httptreemux.ContextMux
	application *newrelic.Application
}

// New creates a new Router to be used in place of httprouter.Router.
func New(app *newrelic.Application) *Router {
	return &Router{
		httptreemux.NewContextMux(),
		app,
	}
}

func txnName(method, path string) string {
	return method + " " + path
}

func (r *Router) handle(method string, path string, original http.HandlerFunc) {
	handle := func(w http.ResponseWriter, req *http.Request) {
		name := txnName(method, path)
		txn := r.application.StartTransaction(name)
		writer := txn.SetWebResponse(w)
		txn.SetWebRequestHTTP(req)

		defer func() {
			txn.End()
		}()

		req = newrelic.RequestWithTransactionContext(req, txn)

		original(writer, req)
	}

	r.ContextMux.Handle(method, path, handle)
}

// DELETE replaces httprouter.Router.DELETE.
func (r *Router) DELETE(path string, h http.HandlerFunc) {
	r.handle(http.MethodDelete, path, h)
}

// GET replaces httprouter.Router.GET.
func (r *Router) GET(path string, h http.HandlerFunc) {
	r.handle(http.MethodGet, path, h)
}

// HEAD replaces httprouter.Router.HEAD.
func (r *Router) HEAD(path string, h http.HandlerFunc) {
	r.handle(http.MethodHead, path, h)
}

// OPTIONS replaces httprouter.Router.OPTIONS.
func (r *Router) OPTIONS(path string, h http.HandlerFunc) {
	r.handle(http.MethodOptions, path, h)
}

// PATCH replaces httprouter.Router.PATCH.
func (r *Router) PATCH(path string, h http.HandlerFunc) {
	r.handle(http.MethodPatch, path, h)
}

// POST replaces httprouter.Router.POST.
func (r *Router) POST(path string, h http.HandlerFunc) {
	r.handle(http.MethodPost, path, h)
}

// PUT replaces httprouter.Router.PUT.
func (r *Router) PUT(path string, h http.HandlerFunc) {
	r.handle(http.MethodPut, path, h)
}

// Handle replaces httprouter.Router.Handle.
func (r *Router) Handle(method, path string, h http.HandlerFunc) {
	r.handle(method, path, h)
}

// Handler replaces httprouter.Router.Handler.
func (r *Router) Handler(method, path string, handler http.Handler) {
	_, h := newrelic.WrapHandle(r.application, txnName(method, path), handler)
	r.ContextMux.Handler(method, path, h)
}

// HandlerFunc replaces httprouter.Router.HandlerFunc.
func (r *Router) HandlerFunc(method, path string, handler http.HandlerFunc) {
	r.Handler(method, path, handler)
}

// ServeHTTP replaces httprouter.Router.ServeHTTP.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	result, _ := r.ContextMux.Lookup(w, req)

	if result.StatusCode == http.StatusNotFound {
		txn := r.application.StartTransaction("NotFound")
		writer := txn.SetWebResponse(w)
		txn.SetWebRequestHTTP(req)

		defer func() {
			txn.End()
		}()

		w = writer
	}

	r.ContextMux.ServeHTTP(w, req)
}
