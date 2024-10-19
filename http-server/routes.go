package httpserver

import (
	"github.com/gorilla/mux"
)

func (cs *ClasslyServer) SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/version", cs.getVersion).Methods("GET")
	r.HandleFunc("/user/{username}", cs.getUserDetails).Methods("GET")
	r.HandleFunc("/signup", cs.signUp).Methods("POST")
	r.HandleFunc("/classes", cs.createClass).Methods("POST")
	r.HandleFunc("/bookings", cs.bookClass).Methods("POST")

	return r
}
