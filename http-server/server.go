package httpserver

import (
	"classly/classly"
	"log"
	"net/http"
	"sync"
)

// TODO: Retrieve the server port from the environment variables
const portNumber = ":8080" // Listen on all network interfaces on port 8080

type ClasslyServer struct {
	classly   *classly.Classly
	waitGroup *sync.WaitGroup
}

func InitializeClasslyServer(classly *classly.Classly) *ClasslyServer {
	cs := ClasslyServer{
		classly:   classly,
		waitGroup: &sync.WaitGroup{},
	}
	routes := cs.SetupRoutes()

	cs.waitGroup.Add(1)
	go func() {
		log.Printf("Server is listening on port %v...", portNumber)
		if err := http.ListenAndServe(portNumber, routes); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	return &cs
}

func (cs *ClasslyServer) Close() {
	log.Println("Shutting down the server...")
	cs.waitGroup.Done()
}
