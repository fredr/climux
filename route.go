package climux

import (
	"strings"
)

type Route struct {
	handler Handler
	path    string
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
