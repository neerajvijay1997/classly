package httpserver

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (cs *ClasslyServer) getVersion(w http.ResponseWriter, r *http.Request) {
	version := cs.classly.GetVersion()
	response := GetVersionResponse{Version: version}
	cs.writeJSONResponse(w, http.StatusOK, response)
}

func (cs *ClasslyServer) signUp(w http.ResponseWriter, r *http.Request) {
	var signUpRequest SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&signUpRequest); err != nil {
		cs.writeErrorResponse(w, http.StatusBadRequest, "Invalid input")
		return
	}

	userName := cs.classly.CreateUser(signUpRequest.Name, signUpRequest.Email)
	response := SignUpResponse{
		Message:  "User registered successfully",
		UserName: userName,
	}
	cs.writeJSONResponse(w, http.StatusCreated, response)
}

func (cs *ClasslyServer) getUserDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userName := params["username"]

	user, exists := cs.classly.GetUserInfo(userName)
	if !exists {
		cs.writeErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}

	response := GetUserDetailsResponse{
		UserName: userName,
		Name:     user.Name,
		Email:    user.Email,
	}
	cs.writeJSONResponse(w, http.StatusOK, response)
}

func (cs *ClasslyServer) createClass(w http.ResponseWriter, r *http.Request) {
	var request CreateClassRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		cs.writeErrorResponse(w, http.StatusBadRequest, "Invalid input")
		return
	}

	classID, err := cs.classly.CreateClass(request.UserName, request.ClassName, request.StartDate, request.EndDate, request.Capacity)
	if err != nil {
		cs.writeErrorResponse(w, http.StatusInternalServerError, "Failed to create class")
		return
	}

	response := CreateClassResponse{
		Message: "Class created successfully",
		ClassID: classID,
	}
	cs.writeJSONResponse(w, http.StatusCreated, response)
}

func (cs *ClasslyServer) bookClass(w http.ResponseWriter, r *http.Request) {
	var request BookClassRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		cs.writeErrorResponse(w, http.StatusBadRequest, "Invalid input")
		return
	}

	classSessionId, err := cs.classly.BookClass(request.UserName, request.ClassID, request.BookingDate)
	if err != nil {
		cs.writeErrorResponse(w, http.StatusInternalServerError, "Failed to book class")
		return
	}

	response := BookClassResponse{
		Message:        "Class booked successfully",
		ClassSessionId: classSessionId,
	}

	cs.writeJSONResponse(w, http.StatusCreated, response)
}

func (cs *ClasslyServer) writeJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		cs.writeErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
}

func (cs *ClasslyServer) writeErrorResponse(w http.ResponseWriter, status int, message string) {
	errorResponse := ErrorResponse{Message: message}
	cs.writeJSONResponse(w, status, errorResponse)
}
