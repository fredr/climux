climux [![GoDoc](https://godoc.org/github.com/fredr/climux?status.svg)](https://godoc.org/github.com/fredr/climux)
======

A mux for the cli in go


### Install
    go get github.com/fredr/climux

### Test
    go test -v

### Use
```go
package main

import (
	"fmt"
	"github.com/fredr/climux"
)

func main() {

	r := climux.NewRouter()
	r.HandleFunc("hello {name}", hello, "says hello to {name}")
	r.HandleFunc("hi {firstname} [lastname]", hi, "says hi to {firstname} with [lastname] if present")
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

func hi(r *climux.Request) {
	if lastName, ok := r.Vars()["lastname"]; ok {
	    fmt.Printf("Hi, %s %s!", r.Vars()["firstname"], lastName)
	} else {
		fmt.Printf("Hi, %s!", r.Vars()["firstname"])
	}
}
```
