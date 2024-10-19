package store

import (
	"classly/types"
)

type Store interface {
	SetUser(user types.User)
	GetUser(userName string) (types.User, bool)
	SetClass(class types.Class)
	GetClass(classId string) (types.Class, bool)
	GetAllClasses() []types.Class
	AddUserToClassSession(sessionId string, userName string)
	AddOfferedClassToUser(userName string, classId string)
}
