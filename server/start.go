package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gouniverse/webserver"
	"github.com/mingrammer/cfmt"
)

const TestingMode = "testing"
const ProductionMode = "production"

// DefaultMode is the default mode for the server.
const DefaultMode = ProductionMode

// Options represents the configuration for the web server.
type Options struct {
	Host    string
	Port    string
	URL     string // optional, displayed in logs
	Handler func(w http.ResponseWriter, r *http.Request)
	Mode    string // optional, default is production, can be development or testing
}

var shutdownChan = make(chan os.Signal, 1)

// StartWebServerbserver starts the web server at the specified host and port and listens
// for incoming requests.
//
// Example:
//
//	StartWebServer(Options{
//	 Host: "localhost",
//	 Port: "8080",
//	 Handler: func(w http.ResponseWriter, r *http.Request) {},
//	 Mode: "production",
//	})
//
// Parameters:
// - none
//
// Returns:
// - none
func Start(options Options) (server *webserver.Server, err error) {
	// Set default mode if not provided
	if options.Mode == "" {
		options.Mode = DefaultMode
	}

	// Create the server address
	addr := options.Host + ":" + options.Port

	// Log server startup
	cfmt.Infoln("Starting server on", addr)
	if options.URL != "" {
		cfmt.Infoln("APP URL:", options.URL)
	}

	// Create a new web server
	server = webserver.New(addr, options.Handler)

	// Register shutdown signals
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGTERM)

	// Start the server in a separate goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if options.Mode == TestingMode {
				cfmt.Errorln(err)
			} else {
				log.Fatal(err)
			}
		}
	}()

	// Wait for a shutdown signal
	cfmt.Infoln("Server is running, press Ctrl+C to stop it.")
	sig := <-shutdownChan
	cfmt.Infoln("Received signal:", sig)
	cfmt.Infoln("Shutting down server...")

	// Shutdown the server
	if err := server.Shutdown(context.Background()); err != nil {
		return nil, err
	}

	if options.Mode != TestingMode {
		os.Exit(0)
	}

	return server, nil
}
