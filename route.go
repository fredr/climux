package climux

import (
	"fmt"
	"strings"
)

// Route holds the information to match against Requests
type Route struct {
	handler     Handler
	path        string
	description string
}

func (r Route) String() string {
	if r.description == "" {
		return r.path
	}

	return fmt.Sprintf("  %s: %s", r.path, r.description)
}

func isParameter(s string) bool {
	return s[:1] == "{" && s[len(s)-1:] == "}"
}

func (r *Route) getVars(req *Request) map[string]string {
	vars := map[string]string{}
	pathArgs := strings.Split(r.path, " ")
	for i, arg := range req.args {
		if isParameter(pathArgs[i]) {
			key := pathArgs[i][1 : len(pathArgs[i])-1]
			vars[key] = arg
		}
	}

	return vars
}

func (r *Route) match(req *Request) bool {
	pathArgs := strings.Split(r.path, " ")
	if len(pathArgs) != len(req.args) {
		return false
	}

	for i, arg := range req.args {
		if !isParameter(pathArgs[i]) {
			if arg != pathArgs[i] {
				return false
			}
		}
	}

	return true
}
