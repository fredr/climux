package climux

type Router struct {
	Routes []Route
}

// HandleFunc adds a Handler to a path
func (r *Router) HandleFunc(path string, fn Handler) {
	route := Route{
		handler: fn,
		path:    path,
	}

	r.Routes = append(r.Routes, route)
}

// NewRouter returns a new Router instance
func NewRouter() *Router {
	return &Router{
		Routes: []Route{},
	}
}
