package httpserver

import (
	"encoding/json"
	"fmt"
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
		cs.writeErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Invalid input in signup request: %v", err))
		return
	}

	userName, err := cs.classly.CreateUser(signUpRequest.Name, signUpRequest.Email)
	if err != nil {
		cs.writeErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create user: %v", err))
		return
	}

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
		cs.writeErrorResponse(w, http.StatusNotFound, fmt.Sprintf("User %v not found", userName))
		return
	}

	response := GetUserDetailsResponse{
		UserName: userName,
		Name:     user.Name,
		Email:    user.Email,
	}
	cs.writeJSONResponse(w, http.StatusOK, response)
}

func (cs *ClasslyServer) getBookedClasses(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userName := params["username"]

	bookedClasses, err := cs.classly.GetBookedClasses(userName)
	if err != nil {
		cs.writeErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get booked classes: %v", err))
		return
	}

	var message string
	if len(bookedClasses) == 0 {
		message = "No booked classes found."
	} else {
		message = "Here are your booked classes."
	}

	response := GetBookedClassesResponse{
		Message:       message,
		BookedClasses: bookedClasses,
	}

	cs.writeJSONResponse(w, http.StatusOK, response)
}

func (cs *ClasslyServer) getAllClasses(w http.ResponseWriter, r *http.Request) {
	classes, err := cs.classly.GetAllClasses()
	if err != nil {
		cs.writeErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get all classes: %v", err))
		return
	}

	var message string
	if len(classes) == 0 {
		message = "No classes are available at the moment."
	} else {
		message = "Here are the available classes."
	}

	response := GetAllClassesResponse{
		Message: message,
		Classes: classes,
	}

	cs.writeJSONResponse(w, http.StatusOK, response)
}

func (cs *ClasslyServer) getClassesStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userName := params["username"]

	classesStatus, err := cs.classly.GetClassesStatus(userName)
	if err != nil {
		cs.writeErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get classes status: %v", err))
		return
	}

	var message string
	if len(classesStatus) == 0 {
		message = "No created class found"
	} else {
		message = "Here are your created classes."
	}

	response := GetClassesStatusResponse{
		Message:       message,
		ClassesStatus: classesStatus,
	}

	cs.writeJSONResponse(w, http.StatusOK, response)
}

func (cs *ClasslyServer) createClass(w http.ResponseWriter, r *http.Request) {
	var request CreateClassRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		cs.writeErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Invalid input in create class request: %v", err))
		return
	}

	classID, err := cs.classly.CreateClass(request.UserName, request.ClassName, request.Description, request.StartDate, request.EndDate, request.Capacity)
	if err != nil {
		cs.writeErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create class: %v", err))
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
		cs.writeErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Invalid input in book class request: %v", err))
		return
	}

	classSessionId, err := cs.classly.BookClass(request.UserName, request.ClassID, request.BookingDate)
	if err != nil {
		cs.writeErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Failed to book class: %v", err))
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
