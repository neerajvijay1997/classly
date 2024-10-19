package main

import (
	"classly/classly"
	httpserver "classly/http-server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// TODO: Initialize store

	classly := classly.InitializeClassly()
	classlyServer := httpserver.InitializeClasslyServer(classly)

	stopChan := make(chan os.Signal, 2)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan // wait for interrupt or terminate signal

	classlyServer.Close()
}
