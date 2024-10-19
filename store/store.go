package store

import (
	"classly/types"
)

type Store interface {
	SetUser(user types.User)
	GetUser(userName string) (types.User, bool)
}
