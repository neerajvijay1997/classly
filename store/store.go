package store

import (
	"classly/utils"
	"time"
)

type Store interface {
	// SetUser saves the specified user to the store.
	// It returns an error if the operation fails.
	SetUser(user utils.User) error

	// GetUser retrieves the user associated with the specified username.
	// It returns the user and a boolean indicating whether the user exists.
	GetUser(userName string) (utils.User, bool)

	// SetClass saves the specified class to the store.
	// It returns an error if the operation fails.
	SetClass(class utils.Class) error

	// GetClass retrieves the class associated with the specified class ID.
	// It returns the class and a boolean indicating whether the class exists.
	GetClass(classId string) (utils.Class, bool)

	// GetAllClasses retrieves all classes available in the store.
	// It returns a slice of Class objects and an error if the retrieval fails.
	GetAllClasses() ([]utils.Class, error)

	// BookClass books a class for a user at the specified booking date.
	// It returns the booking session ID and an error if the booking fails.
	BookClass(userName string, classId string, bookingDate time.Time) (string, error)

	// UpdateUserWithCreatedClass updates the user's record with the ID of the class they created.
	// It returns an error if the operation fails.
	UpdateUserWithCreatedClass(userName string, classId string) error

	// GetBookedClasses retrieves the list of classes booked by the specified user.
	// It returns a slice of BookedClass objects and an error if the retrieval fails.
	GetBookedClasses(userName string) ([]utils.BookedClass, error)

	// GetClassesStatus retrieves the status of classes created by the specified user.
	// It returns a slice of ClassStatus objects and an error if the retrieval fails.
	GetClassesStatus(userName string) ([]utils.ClassStatus, error)
}
