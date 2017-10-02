package web

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Server will create a http.Server from the Go standard library
type Server struct {
	Addr    string        // TCP address to listen on, ":http" if empty
	Handler http.Handler  // handler to invoke, http.DefaultServeMux if nil
	Timeout time.Duration // Timeout to wait for shutdown
}

// ListenAndServe mirrors the function from the Go standard library
// ListenAndServe always returns a non-nil error
func (s *Server) ListenAndServe() error {
	server := http.Server{
		Addr:    s.Addr,
		Handler: s.Handler,
	}

	// We want to use an error channel to block and receive the error
	serverErr := make(chan error, 1)

	// Start the listener.
	go func() {
		serverErr <- server.ListenAndServe()
	}()

	// Listen for an interrupt signal from the OS. Use a buffered
	// channel because of how the signal package is implemented.
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	// If server.ListenAndServe() cannot startup due to errors such as "port in use",
	// it will return error and get stuck waiting for osSignals.
	// This is not ideal because if a server doesn't start, we want to log the error and exit.
	// Unfortunately, we can't select on a waitgroup.
	// The done channel and select statement is used to handle the above case.
	select {
	case err := <-serverErr:
		return err
	case <-osSignals:
		ctx, cancel := context.WithTimeout(context.Background(), s.Timeout)
		defer cancel()

		// Attempt the graceful shutdown by closing the listener and
		// completing all inflight requests.
		var sdErr error
		if err := server.Shutdown(ctx); err != nil {
			sdErr = err
			// Looks like we timedout on the graceful shutdown. Kill it hard.
			if err := server.Close(); err != nil {
				sdErr = err
			}
		}
		if sdErr != nil {
			return sdErr
		}
		// If we're in this select block, we can safely collect the error from this channel
		return <-serverErr
	}
}
