package climux

type Router struct {
	Routes []Route
}

func (r *Router) HandleFunc(path string, fn Handler) {
	route := Route{
		handler: fn,
		path:    path,
	}

	r.Routes = append(r.Routes, route)
}

func NewRouter() *Router {
	return &Router{
		Routes: []Route{},
	}
}
