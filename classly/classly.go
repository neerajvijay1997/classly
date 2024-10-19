package classly

import (
	"classly/store"
	"classly/types"
)

type Classly struct {
	store store.Store
}

func InitializeClassly(store store.Store) *Classly {
	return &Classly{
		store,
	}
}

func (cly *Classly) CreateUser(userName, name string, email string) {
	// TODO: Generate unique usernames on the server side
	user := types.User{
		UserName: userName,
		Name:     name,
		Email:    email,
	}

	cly.store.SetUser(user)
}

func (cly *Classly) GetUserInfo(userName string) (types.User, bool) {
	user, ok := cly.store.GetUser(userName)
	return user, ok
}
