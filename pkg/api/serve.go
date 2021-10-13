package api

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (app *Application) Serve() error {

	// Setup a non-default HTTP server
	srv := &http.Server{
		Addr:         "localhost:8000",
		Handler:      app.GetRoutes(),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	// Create a channel for shutdown errors
	shutdownError := make(chan error)

	// srv.ListenAndServe is blocking so do our work outside this.
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		// Block until we have an event SIGINT or SIGTERM
		s := <-quit
		app.InfoLog.Println("shutting down server: ", s)

		// Create a cancellation context to permit graceful shutdown of goroutines.
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// svr.Shutdown is graceful shutdown
		err := srv.Shutdown(ctx)
		if err != nil {
			// Encountered an error when trying to gracefully shutdown so capture that error
			shutdownError <- err
		}

		// We didn't encounter an error as we shut down gracefully without error
		shutdownError <- nil

	}()

	app.InfoLog.Println("starting server on port 8000")

	err := srv.ListenAndServe()

	// ErrServerClosed is the desired error, it implies a graceful shutdown. If we don't get this error, tell us why.
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	// As Above, we received an error when we tried to gracefully shutdown.
	err = <-shutdownError
	if err != nil {
		return err
	}

	app.InfoLog.Println("stopped server")
	return nil
}
