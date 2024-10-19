package store

import (
	"classly/types"
)

type MemStore struct {
	users map[string]types.User
}

func NewMemStore() *MemStore {
	return &MemStore{
		users: make(map[string]types.User),
	}
}

func (ms *MemStore) SetUser(user types.User) {
	ms.users[user.UserName] = user
}

func (ms *MemStore) GetUser(userName string) (types.User, bool) {
	user, ok := ms.users[userName]
	return user, ok
}
