package climux

import (
	"testing"
)

func TestRouteMatchFail(t *testing.T) {
	req := &Request{
		args: []string{
			"some",
			"damn",
			"path",
		},
	}
	route := Route{
		path: "an other path",
	}

	if route.match(req) {
		t.Errorf("Expected %v to not match %v", req, route)
	}
}

func TestSimpleRouteMatch(t *testing.T) {
	req := &Request{
		args: []string{
			"some",
			"damn",
			"path",
		},
	}
	route := Route{
		path: "some damn path",
	}

	if !route.match(req) {
		t.Errorf("Expected %v to match %v", req, route)
	}
}

func TestParameterizedRouteMatch(t *testing.T) {
	req := &Request{
		args: []string{
			"do",
			"this",
			"then",
			"that",
		},
	}
	route := Route{
		path: "do {a} then {b}",
	}

	if !route.match(req) {
		t.Errorf("Expected %v to match %v", req, route)
	}
}

func TestGetVars(t *testing.T) {
	req := &Request{
		args: []string{
			"do",
			"this",
			"then",
			"that",
		},
	}
	route := Route{
		path: "do {a} then {b}",
	}

	vars := route.getVars(req)
	if vars["a"] != "this" {
		t.Errorf("Expected %q to match %q", "a", "this")
	}
	if vars["b"] != "that" {
		t.Errorf("Expected %q to match %q", "b", "that")
	}
}
