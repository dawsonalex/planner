package main

import (
	"context"
	"github.com/dawsonalex/golang-cli/html"
	"log"
	"net/http"
	"os"
	"os/signal"
)

var ()

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		err := html.Dashboard(w, html.DashboardParams{Greeting: "Good day!"})
		if err != nil {
			log.Printf("error displaying dashboard: %+v", err)
		}
	})
	srv := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Wait here until SIGINT received, then exec callback function
	// to gracefully shutdown.
	awaitInterrupt(func(done chan bool) {
		if err := srv.Shutdown(context.Background()); err != nil {
			panic(err)
		}
		done <- true
	})
}
func awaitInterrupt(onInterrupt func(chan bool)) {
	done := make(chan bool)
	go func() {
		// Wait for SIGINT to stop services.
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, os.Interrupt)
		defer signal.Stop(sigchan)
		<-sigchan

		go onInterrupt(done)
	}()

	<-done
}
