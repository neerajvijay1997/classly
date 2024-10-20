package store

import (
	"classly/utils"
	"time"
)

type Store interface {
	SetUser(user utils.User)
	GetUser(userName string) (utils.User, bool)
	SetClass(class utils.Class)
	GetClass(classId string) (utils.Class, bool)
	GetAllClasses() []utils.Class
	BookClass(userName string, classId string, bookingDate time.Time) (string, error)
	UpdateUserWithCreatedClass(userName string, classId string)
	GetBookedClasses(userName string) ([]utils.BookedClass, error)
	GetClassesStatus(userName string) ([]utils.ClassStatus, error)
}
