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

## Object

The object package provides a flexible and thread-safe implementation for managing properties and serializable objects.
It offers interfaces and implementations for property storage and JSON serialization.

For more information, see the [object/README.md](object/README.md) file.


## Router

The router package provides HTTP routing functionality with middleware support for the Dracory framework.
It enables building flexible and powerful web applications with clean routing patterns.

For more information, see the [router/README.md](router/README.md) file.

## Server

The server package provides web server functionality for the Dracory framework.
It offers a simple and configurable way to create and manage HTTP servers.

For more information, see the [server/README.md](server/README.md) file.
