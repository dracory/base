# CFMT

A simple and modern colored formatting package for Go console output.

## Features

- Colored console output with ANSI escape codes
- Support for different message types (Info, Success, Warning, Error)
- Formatted string support (Printf-style)
- Customizable output writer
- Thread-safe operations

## Installation

```bash
go get github.com/dracory/base/cfmt
```

## Usage

```go
package main

import "github.com/dracory/base/cfmt"

func main() {
    // Basic usage
    cfmt.Info("This is an info message")
    cfmt.Success("Operation completed successfully")
    cfmt.Warning("This is a warning message")
    cfmt.Error("An error occurred")

    // With newline
    cfmt.Infoln("This is an info message with newline")
    cfmt.Errorln("This is an error message with newline")

    // Formatted output
    cfmt.Infof("Hello, %s!", "World")
    cfmt.Errorf("Error: %s", "Something went wrong")

    // Custom output writer
    var buf bytes.Buffer
    cfmt.SetOutput(&buf)
    cfmt.Info("This will be written to the buffer")
}
```

## Available Functions

### Basic Print Functions
- `Print(color string, a ...interface{})`
- `Println(color string, a ...interface{})`
- `Printf(color string, format string, a ...interface{})`

### Info Functions (Blue)
- `Info(a ...interface{})`
- `Infoln(a ...interface{})`
- `Infof(format string, a ...interface{})`

### Success Functions (Green)
- `Success(a ...interface{})`
- `Successln(a ...interface{})`
- `Successf(format string, a ...interface{})`

### Warning Functions (Yellow)
- `Warning(a ...interface{})`
- `Warningln(a ...interface{})`
- `Warningf(format string, a ...interface{})`

### Error Functions (Red)
- `Error(a ...interface{})`
- `Errorln(a ...interface{})`
- `Errorf(format string, a ...interface{})`

### Configuration
- `SetOutput(w io.Writer)`

## Color Codes

The package provides the following color constants:
- `Reset`
- `Bold`
- `Red`, `Green`, `Yellow`, `Blue`, `Magenta`, `Cyan`, `White`
- `BoldRed`, `BoldGreen`, `BoldYellow`, `BoldBlue`

## License

This package is part of the dracory/base project and is licensed under the same terms. 