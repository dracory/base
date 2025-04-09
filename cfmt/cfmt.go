package cfmt

import (
	"fmt"
	"io"
	"os"
)

// Colors for terminal output
const (
	Reset      = "\033[0m"
	Bold       = "\033[1m"
	Red        = "\033[31m"
	Green      = "\033[32m"
	Yellow     = "\033[33m"
	Blue       = "\033[34m"
	Magenta    = "\033[35m"
	Cyan       = "\033[36m"
	White      = "\033[37m"
	BoldRed    = Bold + Red
	BoldGreen  = Bold + Green
	BoldYellow = Bold + Yellow
	BoldBlue   = Bold + Blue
)

var (
	// Default output writer
	output io.Writer = os.Stdout
)

// SetOutput sets the output writer for the package
func SetOutput(w io.Writer) {
	output = w
}

// Print prints the arguments with the given color
func Print(color string, a ...interface{}) {
	fmt.Fprint(output, color)
	fmt.Fprint(output, a...)
	fmt.Fprint(output, Reset)
}

// Println prints the arguments with the given color and adds a newline
func Println(color string, a ...interface{}) {
	fmt.Fprint(output, color)
	fmt.Fprintln(output, a...)
	fmt.Fprint(output, Reset)
}

// Printf prints a formatted string with the given color
func Printf(color string, format string, a ...interface{}) {
	fmt.Fprint(output, color)
	fmt.Fprintf(output, format, a...)
	fmt.Fprint(output, Reset)
}

// Info prints information in blue
func Info(a ...interface{}) {
	Print(BoldBlue, a...)
}

// Infoln prints information in blue with a newline
func Infoln(a ...interface{}) {
	Println(BoldBlue, a...)
}

// Infof prints formatted information in blue
func Infof(format string, a ...interface{}) {
	Printf(BoldBlue, format, a...)
}

// Success prints success message in green
func Success(a ...interface{}) {
	Print(BoldGreen, a...)
}

// Successln prints success message in green with a newline
func Successln(a ...interface{}) {
	Println(BoldGreen, a...)
}

// Successf prints formatted success message in green
func Successf(format string, a ...interface{}) {
	Printf(BoldGreen, format, a...)
}

// Warning prints warning message in yellow
func Warning(a ...interface{}) {
	Print(BoldYellow, a...)
}

// Warningln prints warning message in yellow with a newline
func Warningln(a ...interface{}) {
	Println(BoldYellow, a...)
}

// Warningf prints formatted warning message in yellow
func Warningf(format string, a ...interface{}) {
	Printf(BoldYellow, format, a...)
}

// Error prints error message in red
func Error(a ...interface{}) {
	Print(BoldRed, a...)
}

// Errorln prints error message in red with a newline
func Errorln(a ...interface{}) {
	Println(BoldRed, a...)
}

// Errorf prints formatted error message in red
func Errorf(format string, a ...interface{}) {
	Printf(BoldRed, format, a...)
}
