package httpserver

import (
	"classly/classly"
	"log"
	"net/http"
	"sync"
)

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
		// TODO: Retrieve the server port from the environment variables
		log.Println("Server is listening on port 8080...")
		if err := http.ListenAndServe(":8080", routes); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	return &cs
}

func (cs *ClasslyServer) Close() {
	log.Println("Shutting down the server...")
	cs.waitGroup.Done()
}
