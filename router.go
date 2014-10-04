package climux

// Router holds all routes
type Router struct {
	Routes          []Route
	NotFoundHandler Handler
}

// HandleFunc adds a Handler to a path
func (r *Router) HandleFunc(path string, fn Handler, description string) {
	route := Route{
		handler:     fn,
		path:        path,
		description: description,
	}

	r.Routes = append(r.Routes, route)
}

// NewRouter returns a new Router instance
func NewRouter() *Router {
	return &Router{
		Routes: []Route{},
	}
}
