package hayane

import "net/http"

type Haya struct {
	Router *Router
}

type Middleware interface {
	Handle(next http.Handler) http.Handler
}

func Create() *Haya {
	return &Haya{
		Router: &Router{},
	}
}

func (h *Haya) Add(middleware interface{}) {
	switch mw := middleware.(type) {
	case func(http.Handler) http.Handler:
		h.Router.AddMiddleware(mw)
	case Middleware:
		h.Router.AddMiddleware(mw.Handle)
	default:
		panic("kyaa!! Unsupported middleware type!")
	}
}

func (h *Haya) GET(path string, handler func(*Context)) {
	h.Router.AddRoute("GET", path, handler)
}

func (h *Haya) POST(path string, handler func(*Context)) {
	h.Router.AddRoute("POST", path, handler)
}

func (h *Haya) Run(addr string) {
	http.ListenAndServe(addr, h.Router)
}
