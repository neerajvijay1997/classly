package store

import (
	"classly/types"
	"time"
)

type Store interface {
	SetUser(user types.User)
	GetUser(userName string) (types.User, bool)
	SetClass(class types.Class)
	GetClass(classId string) (types.Class, bool)
	GetAllClasses() []types.Class
	BookClass(userName string, classId string, bookingDate time.Time) (string, error)
	UpdateUserWithCreatedClass(userName string, classId string)
	GetBookedClasses(userName string) ([]types.BookedClass, error)
	GetClassesStatus(userName string) ([]types.ClassStatus, error)
}
