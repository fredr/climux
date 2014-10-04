climux
======

A mux for the cli in Go


### Install
    go get github.com/fredr/climux

### Test
    go test -v

### Use
```
package main

import (
	"fmt"
	"github.com/fredr/climux"
)

func main() {

	r := climux.NewRouter()
	r.HandleFunc("hello {name}", hello, "says hello to {name}")
	r.HandleFunc("help", help(r), "shows help")
	r.NotFoundHandler = help(r)

	climux.Handle(r)
}

func help(router *climux.Router) climux.Handler {
	return func(r *climux.Request) {
		fmt.Println("commands:")
		for _, route := range router.Routes {
			fmt.Println(route)
		}
	}
}

func hello(r *climux.Request) {
	fmt.Printf("Hello, %s!", r.Vars()["name"])
}

```
