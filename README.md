# bassnectar

Basic server implementation that I like for building out my microservices at this time.

## Installation

```bash
go get github.com/noiseunion/bassnectar
```

or if you are using `go mod` just import it into your project.

## Usage

The server can be pulled into your project and used to define an HTTP server.

```golang
package main

import (
	"net/http"

	"github.com/go-chi/cors"
	"github.com/noiseunion/bassnectar/pkg/env"
	"github.com/noiseunion/bassnectar/pkg/routing"
	"github.com/noiseunion/bassnectar/pkg/server"
)

var instance *server.Instance
var config server.Config

// This file is pretty much just an example of how to use the packages
// in this library.
func init() {
	env.LoadDotEnv()

	// Load our server config
	config = server.Config{
		RouteBuilder: registerRoutes(),
		CORS: cors.Options{
			AllowedOrigins: []string{"*", "localhost"},
		},
	}

	config.LoadEnv()
}

func registerRoutes() *routing.Builder {
	b := routing.NewRouteBuilder()

	b.Register(http.MethodGet, "/foo", handleFoo)

	return b
}

func handleFoo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("FOO BAR"))
}

func main() {
	i := server.New(&config)
	i.Start()
}
```

An example can be found under `_examples`.  You can run the example locally and test it out.