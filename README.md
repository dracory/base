# base <a href="https://gitpod.io/#https://github.com/dracory/base" style="float:right:"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

[![Tests Status](https://github.com/dracory/base/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/dracory/base/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/dracory/base)](https://goreportcard.com/report/github.com/dracory/base)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/dracory/base)](https://pkg.go.dev/github.com/dracory/base)

## License

This project is dual-licensed under the following terms:

- For non-commercial use, you may choose either the GNU Affero General Public License v3.0 (AGPLv3) *or* a separate commercial license (see below). You can find a copy of the AGPLv3 at: https://www.gnu.org/licenses/agpl-3.0.txt

- For commercial use, a separate commercial license is required. Commercial licenses are available for various use cases. Please contact me via my [contact page](https://lesichkov.co.uk/contact) to obtain a commercial license.

## Installation

```
go get github.com/dracory/base
```


## About Dracory

The Dracory project is a Go framework that provides various utilities, including:

*   BBCode to HTML conversion
*   Slice manipulation
*   Database interaction
*   Error handling and validation
*   Image manipulation
*   URL downloading
*   QR code generation
*   HTTP request handling
*   Timezone conversion
*   Date and datetime validation
*   Web server functionality
*   Command line functionality
*   HTTP routing with middleware support

## Environment Variables

The Dracory framework provides easy access to environment variables
using the `env` package.

For information on environment variables, see the [env/README.md](env/README.md) file.

## Database

The database package provides database interaction functionalities for the Dracory framework.
It offers a set of tools for interacting with various database systems.

For more information, see the [database/README.md](database/README.md) file.

## Router

The router package provides a flexible and intuitive way to define HTTP routes in your application. It supports method chaining and includes shortcut methods for common HTTP methods.

### Basic Usage

```go
package main

import (
	"net/http"

	"github.com/dracory/base/router"
)

func main() {
	// Create a new router
	r := router.NewRouter()

	// Define a handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}

	// Create routes using shortcut methods
	r.AddRoute(router.Get("/hello", handler))
	r.AddRoute(router.Post("/submit", handler))
	r.AddRoute(router.Put("/update", handler))
	r.AddRoute(router.Delete("/remove", handler))

	// You can also use the traditional method chaining approach
	route := router.NewRoute().
		SetMethod(http.MethodGet).
		SetPath("/custom").
		SetHandler(handler).
		SetName("custom-route")

	// Add middleware to routes
	middleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Do something before the handler
			next.ServeHTTP(w, r)
			// Do something after the handler
		})
	}

	r.AddBeforeMiddlewares([]router.Middleware{middleware})

	// Start the server
	http.ListenAndServe(":8080", r)
}
```

### Available Shortcut Methods

The router package provides the following shortcut methods for creating routes:

- `Get(path string, handler Handler) RouteInterface` - Creates a GET route
- `Post(path string, handler Handler) RouteInterface` - Creates a POST route
- `Put(path string, handler Handler) RouteInterface` - Creates a PUT route
- `Delete(path string, handler Handler) RouteInterface` - Creates a DELETE route

These shortcut methods automatically set the HTTP method, path, and handler for the route, making your code more concise and readable.

### Method Chaining

All route methods support method chaining, allowing you to fluently configure your routes:

```go
route := router.NewRoute().
	SetMethod(http.MethodGet).
	SetPath("/users").
	SetName("users-list").
	SetHandler(userHandler).
	AddBeforeMiddlewares([]router.Middleware{authMiddleware}).
	AddAfterMiddlewares([]router.Middleware{logMiddleware})
```

This approach gives you full control over route configuration when needed.

For more information, see the [router/README.md](router/README.md) file.

## Running a Web Server

Here's an example of how to run a web server using Dracory:

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/dracory/base/server"
)

func main() {
	// Define the handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Dracory!")
	}

	// Define the server options
	options := server.Options{
		Host:    "localhost",
		Port:    "8080",
		Handler: handler,
	}

	// Start the server
	_, err := server.Start(options)
	if err != nil {
		panic(err)
	}
}
```
