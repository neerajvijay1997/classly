package httpserver

import (
	"github.com/gorilla/mux"
)

func (cs *ClasslyServer) SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/version", cs.getVersion).Methods("GET")
	r.HandleFunc("/user/{username}", cs.getUserDetails).Methods("GET")
	r.HandleFunc("/all-classes", cs.getAllClasses).Methods("GET")
	r.HandleFunc("/booked-classes/{username}", cs.getBookedClasses).Methods("GET")
	r.HandleFunc("/classes-status/{username}", cs.getClassesStatus).Methods("GET")
	r.HandleFunc("/signup", cs.signUp).Methods("POST")
	r.HandleFunc("/classes", cs.createClass).Methods("POST")
	r.HandleFunc("/bookings", cs.bookClass).Methods("POST")

	return r
}
