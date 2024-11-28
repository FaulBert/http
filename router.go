package hayane

import (
	"net/http"
)

type route struct {
	method  string
	path    string
	handler http.HandlerFunc
}

type Router struct {
	routes     []route
	middleware []func(http.Handler) http.Handler
}

func (r *Router) AddMiddleware(mw func(http.Handler) http.Handler) {
	r.middleware = append(r.middleware, mw)
}

func (r *Router) AddRoute(method, path string, handler func(*Context)) {
	final := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := &Context{Writer: w, Request: req}
		handler(ctx)
	})

	// TODO: FIX middleware!!!
	for i := len(r.middleware) - 1; i >= 0; i-- {
		//final = r.middleware[i](final

		final = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			r.middleware[i](final).ServeHTTP(w, req)
		})
	}

	r.routes = append(r.routes, route{method, path, final.ServeHTTP})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if req.Method == route.method && req.URL.Path == route.path {
			route.handler(w, req)
			return
		}
	}

	http.NotFound(w, req)
}
