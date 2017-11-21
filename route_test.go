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

func TestOptionalParameterRouteMatch(t *testing.T) {
	route := Route{
		path: "{a} [b]",
	}

	req := &Request{
		args: []string{
			"A",
			"B",
		},
	}
	if !route.match(req) {
		t.Errorf("Expected %v to match %v", req, route)
	}

	req = &Request{
		args: []string{
			"A",
		},
	}
	if !route.match(req) {
		t.Errorf("Expected %v to match %v", req, route)
	}

	req = &Request{
		args: []string{
			"A",
			"B",
			"C",
		},
	}
	if route.match(req) {
		t.Errorf("Expected %v to not match %v", req, route)
	}

}

func TestOptionalParameterGetVars(t *testing.T) {
	route := Route{
		path: "{a} [b]",
	}

	req := &Request{
		args: []string{
			"A",
			"B",
		},
	}
	vars := route.getVars(req)
	if vars["a"] != "A" {
		t.Errorf("Expected %q to match %q", "a", "A")
	}
	if vars["b"] != "B" {
		t.Errorf("Expected %q to match %q", "b", "B")
	}

	req = &Request{
		args: []string{
			"A",
		},
	}
	vars = route.getVars(req)
	if vars["a"] != "A" {
		t.Errorf("Expected %q to match %q", "a", "A")
	}
	if vars["b"] != "" {
		t.Errorf("Expected %q to match %q", "b", "")
	}

}
