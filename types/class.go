package types

import "time"

type Class struct {
	Id                    string
	ClassName             string
	ClassProviderUserName string
	StartDate             time.Time
	EndDate               time.Time
	Capacity              uint32
}

// TODO: Add constructor for Class struct

