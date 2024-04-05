package http

import "net/http"

type HandlerFunc func(Context) error

type Router struct {
	routes map[string]map[string]HandlerFunc
}

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

func New() *Router {
	return &Router{
		routes: make(map[string]map[string]HandlerFunc),
	}
}

func (r *Router) GET(path string, handler HandlerFunc) {
	r.addRoute("GET", path, handler)
}

func (r *Router) POST(path string, handler HandlerFunc) {
	r.addRoute("POST", path, handler)
}

func (r *Router) addRoute(method, path string, handler HandlerFunc) {
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]HandlerFunc)
	}
	r.routes[method][path] = handler
}

func (ctx Context) String(response string) error {
	ctx.ResponseWriter.WriteHeader(http.StatusOK)
	_, err := ctx.ResponseWriter.Write([]byte(response))
	return err
}

func (r *Router) Serve(port string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		method := req.Method
		path := req.URL.Path

		if handlers, ok := r.routes[method]; ok {
			if handler, ok := handlers[path]; ok {
				ctx := Context{
					ResponseWriter: w,
					Request:        req,
				}
				if err := handler(ctx); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}
		}

		http.NotFound(w, req)
	})

	return http.ListenAndServe(port, nil)
}
