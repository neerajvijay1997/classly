package main

import (
	"os"
	"os/signal"
	"syscall"

	"classly/classly"
	httpserver "classly/http-server"
	"classly/store"
)

func main() {
	// Initialize Classly instance
	memStore := store.NewMemStore()
	classly := classly.InitializeClassly(memStore)

	// Initialize the HTTP server to serve Classly API
	classlyServer := httpserver.InitializeClasslyServer(classly)

	// Setup a channel to listen for interrupt or terminate signals
	stopChan := make(chan os.Signal, 2)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan // wait for interrupt or terminate signal

	// Close the Classly server gracefully on signal reception
	classlyServer.Close()
}
