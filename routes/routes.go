package routes

import "github.com/gorilla/mux"

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", getVersion).Methods("GET")

	return r
}
