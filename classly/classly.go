package classly

import (
	"classly/store"
	"classly/utils"
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Classly struct {
	store store.Store
}

func InitializeClassly(store store.Store) *Classly {
	return &Classly{
		store,
	}
}

func (cly *Classly) CreateUser(name string, email string) string {
	userName, _ := gonanoid.New(5)

	// TODO: Generate unique usernames on the server side
	user := utils.User{
		UserName:        userName,
		Name:            name,
		Email:           email,
		BookedClasses:   make(utils.BookedClassesMap),
		CreatedClassIds: make([]string, 0),
	}

	cly.store.SetUser(user)
	return userName
}

func (cly *Classly) GetUserInfo(userName string) (utils.User, bool) {
	user, ok := cly.store.GetUser(userName)
	return user, ok
}

func (cly *Classly) CreateClass(userName, className string, startDateStr, endDateStr string, capacity uint32) (string, error) {
	// TODO: Check if the user exists

	startDate, err := utils.ParseTime(startDateStr)
	if err != nil {
		return "", fmt.Errorf("invalid start date format")
	}

	endDate, err := utils.ParseTime(endDateStr)
	if err != nil {
		return "", fmt.Errorf("invalid end date format")
	}

	if endDate.Before(startDate) {
		return "", fmt.Errorf("end date must be after start date")
	}

	classId, err := gonanoid.New(5)
	if err != nil {
		return "", fmt.Errorf("error generating nano id: %w", err)
	}

	class := utils.Class{
		Id:                    classId,
		ClassName:             className,
		ClassProviderUserName: userName,
		Capacity:              capacity,
		StartDate:             startDate,
		EndDate:               endDate,
	}

	cly.store.SetClass(class)
	cly.store.UpdateUserWithCreatedClass(userName, classId)

	return classId, nil
}

func (cly *Classly) GetAllClasses() ([]utils.Class, error) {
	return cly.store.GetAllClasses(), nil
}

func (cly *Classly) GetClassesStatus(userName string) ([]utils.ClassStatus, error) {
	// TODO: If class offered empty return empty arrray or error
	return cly.store.GetClassesStatus(userName)
}

func (cly *Classly) GetBookedClasses(userName string) ([]utils.BookedClass, error) {
	// TODO: If booked class empty return empty arrray
	return cly.store.GetBookedClasses(userName)
}

func (cly *Classly) BookClass(userName string, classId string, bookingDateStr string) (string, error) {
	// TODO: Validate params
	// TODO: Class provider cannot be class member

	bookingDate, err := utils.ParseTime(bookingDateStr)
	if err != nil {
		return "", fmt.Errorf("invalid booking date format")
	}

	classSessionId, _ := cly.store.BookClass(userName, classId, bookingDate)

	return classSessionId, nil
}

func (cly *Classly) GetVersion() string {
	return "Classly-v0.1.0"
}
