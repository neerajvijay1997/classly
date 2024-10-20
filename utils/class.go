package utils

import "time"

type Class struct {
	Id        string `json:"id"`
	ClassName string `json:"class_name"`
	// TODO: Handle
	Description           string    `json:"description"`
	ClassProviderUserName string    `json:"class_provider_user_name"`
	StartDate             time.Time `json:"start_date"`
	EndDate               time.Time `json:"end_date"`
	Capacity              uint32    `json:"capacity"`
}

// TODO: Add constructor for Class struct

type BookedClass struct {
	Class
	Sessions []time.Time
}

type ClassSessionsMap map[time.Time][]User

type ClassStatus struct {
	Class
	Sessions ClassSessionsMap
}
