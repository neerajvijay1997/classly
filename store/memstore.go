package store

import (
	"classly/types"
)

type MemStore struct {
	users         map[string]types.User
	classes       map[string]types.Class
	classSessions map[string][]string
}

func NewMemStore() *MemStore {
	return &MemStore{
		users:         make(map[string]types.User),
		classes:       make(map[string]types.Class),
		classSessions: make(map[string][]string),
	}
}

func (ms *MemStore) SetUser(user types.User) {
	ms.users[user.UserName] = user
}

func (ms *MemStore) updateUserForClassSession(userName string, sessionId string) {
	user := ms.users[userName]
	user.BookedClasses = append(user.BookedClasses, sessionId)
	ms.users[userName] = user
}

func (ms *MemStore) AddOfferedClassToUser(userName string, classId string) {
	user := ms.users[userName]
	user.OfferedClasses = append(user.BookedClasses, classId)
	ms.users[userName] = user
}

func (ms *MemStore) GetUser(userName string) (types.User, bool) {
	user, ok := ms.users[userName]
	return user, ok
}

func (ms *MemStore) SetClass(class types.Class) {
	ms.classes[class.Id] = class
}

func (ms *MemStore) GetClass(classId string) (types.Class, bool) {
	class, ok := ms.classes[classId]
	return class, ok
}

func (ms *MemStore) GetAllClasses() []types.Class {
	allClasses := make([]types.Class, 0, len(ms.classes))
	for _, class := range ms.classes {
		allClasses = append(allClasses, class)
	}
	return allClasses
}

func (ms *MemStore) AddUserToClassSession(classSessionId string, userName string) {
	session, ok := ms.classSessions[classSessionId]
	if ok {
		session = append(session, userName)
	} else {
		session = []string{userName}
	}

	ms.classSessions[classSessionId] = session

	ms.updateUserForClassSession(userName, classSessionId)
}
