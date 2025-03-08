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

## Investigation Summary

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

The BBCode tests are working fine.

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
