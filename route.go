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
		return fmt.Sprintf("  %s", r.path)
	}

	return fmt.Sprintf("  %s: %s", r.path, r.description)
}

func isRequiredParameter(s string) bool {
	return s[:1] == "{" && s[len(s)-1:] == "}"
}

func isOptionalParameter(s string) bool {
	return s[:1] == "[" && s[len(s)-1:] == "]"
}

func isParameter(s string) bool {
	return isRequiredParameter(s) || isOptionalParameter(s)
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
	if len(req.args) > len(pathArgs) {
		return false
	}

	for i := 0; i < len(pathArgs); i++ {
		if i >= len(req.args) {
			if !isOptionalParameter(pathArgs[i]) {
				return false
			}
		} else {
			if !isParameter(pathArgs[i]) {
				if req.args[i] != pathArgs[i] {
					return false
				}
			}
		}
	}

	return true
}
