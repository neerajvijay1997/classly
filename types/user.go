package types

import (
	"time"
)

type User struct {
	UserName string `json:"user_name"`
	Name     string `json:"name"`
	Email    string `json:"email"`

	BookedClasses   BookedClassesMap `json:"-"`
	CreatedClassIds []string         `json:"-"`
}

// TODO: Add constructor for User struct

type BookedClassesMap map[ClassId][]time.Time
type ClassId = string
