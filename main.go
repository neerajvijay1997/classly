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
	memStore := store.NewMemStore()
	classly := classly.InitializeClassly(memStore)
	classlyServer := httpserver.InitializeClasslyServer(classly)

	stopChan := make(chan os.Signal, 2)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan // wait for interrupt or terminate signal

	classlyServer.Close()
}
