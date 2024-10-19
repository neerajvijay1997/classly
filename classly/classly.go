package classly

import (
	"classly/store"
	"classly/types"
	"time"
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

func (cly *Classly) CreateClass(userName, className string, startDateStr, endDateStr string, capacity uint32) {
	// TODO: Check if the user exists
	startDate, _ := ParseTime(startDateStr)
	endDate, _ := ParseTime(endDateStr)

	class := types.Class{
		Id:                    "",
		ClassName:             className,
		ClassProviderUserName: userName,
		Capacity:              capacity,
		StartDate:             startDate,
		EndDate:               endDate,
	}

	cly.store.SetClass(class)
}

func (cly *Classly) GetAllClasses() []types.Class {
	return cly.store.GetAllClasses()
}

// TODO: Move to utils package
// ParseTime converts a string to time.Time based on the provided layout (current layout: "2006-01-02 15:04:05")
func ParseTime(timeStr string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}
