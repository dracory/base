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

type Options struct {
	Host    string
	Port    string
	URL     string // optional, if you want to be displayed in the logs
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
	addr := options.Host + ":" + options.Port
	cfmt.Infoln("Starting server on " + options.Host + ":" + options.Port + " ...")
	if options.URL != "" {
		cfmt.Infoln("APP URL: " + options.URL + " ...")
	}

	server = webserver.New(addr, options.Handler)

	// Register the shutdown signal
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGTERM)

	// Start the server in a separate goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if options.Mode == `testing` {
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
	server.Shutdown(context.Background())
	if options.Mode != `testing` {
		os.Exit(0)
	}

	return server, nil
}
