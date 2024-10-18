package main

import (
	"classly/routes"
	"fmt"
	"net/http"
)

func main() {

	r := routes.SetupRoutes()

	// TODO: Retrieve the server port from the environment variables
	fmt.Println("Server is listening on port 8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
