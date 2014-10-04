package climux

import (
	"flag"
	"strings"
)

type Handler func(*Request)

// Handle builds a request and tries to find a matching route in the router
func Handle(r *Router) {
	if !flag.Parsed() {
		flag.Parse()
	}

	req := &Request{
		args: flag.Args(),
	}

	for _, route := range r.Routes {
		if route.match(req) {
			req.vars = route.getVars(req)
			route.handler(req)
			return
		}
	}
	if r.NotFoundHandler != nil {
		r.NotFoundHandler(req)
	}
}

type Request struct {
	args []string
	vars map[string]string
}

// Vars returns all route variables from the request
func (r Request) Vars() map[string]string {
	return r.vars
}

func (r Request) String() string {
	return strings.Join(r.args, " ")
}
