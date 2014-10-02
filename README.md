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
	r.HandleFunc("hello {name}", hello)

	climux.Handle(r)
}

func hello(r *climux.Request) {
	fmt.Printf("Hello, %s!", r.Vars()["name"])
}

```
