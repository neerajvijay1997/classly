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

// InitializeClassly creates a new instance of Classly with the provided store.
func InitializeClassly(store store.Store) *Classly {
	return &Classly{
		store,
	}
}

// CreateUser generates a unique username and creates a new user with the provided name and email.
func (cly *Classly) CreateUser(name string, email string) (string, error) {
	userName, _ := gonanoid.New(5)

	user := utils.User{
		UserName:        userName,
		Name:            name,
		Email:           email,
		BookedClasses:   make(utils.BookedClassesMap),
		CreatedClassIds: make([]string, 0),
	}

	err := cly.store.SetUser(user)
	if err != nil {
		return "", err
	}

	return userName, nil
}

// GetUserInfo retrieves the user information for the specified username.
func (cly *Classly) GetUserInfo(userName string) (utils.User, bool) {
	user, ok := cly.store.GetUser(userName)
	return user, ok
}

// CreateClass allows a user to create a new class with the specified details.
func (cly *Classly) CreateClass(userName, className string, startDateStr, endDateStr string, capacity uint32) (string, error) {
	_, exist := cly.store.GetUser(userName)
	if !exist {
		return "", fmt.Errorf("user not found")
	}

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

	err = cly.store.SetClass(class)
	if err != nil {
		return "", fmt.Errorf("error storing class: %w", err)
	}

	err = cly.store.UpdateUserWithCreatedClass(userName, classId)
	if err != nil {
		return "", fmt.Errorf("error updating user with created class: %w", err)
	}

	return classId, nil
}

// GetAllClasses retrieves all currently available classes from the store.
func (cly *Classly) GetAllClasses() ([]utils.Class, error) {
	return cly.store.GetAllClasses()
}

// GetClassesStatus retrieves the status of classes created by the specified user.
func (cly *Classly) GetClassesStatus(userName string) ([]utils.ClassStatus, error) {
	return cly.store.GetClassesStatus(userName)
}

// GetBookedClasses retrieves the list of classes booked by the specified user.
func (cly *Classly) GetBookedClasses(userName string) ([]utils.BookedClass, error) {
	return cly.store.GetBookedClasses(userName)
}

// BookClass allows a user to book a specific class on a given date.
func (cly *Classly) BookClass(userName string, classId string, bookingDateStr string) (string, error) {
	_, userExist := cly.store.GetUser(userName)
	if !userExist {
		return "", fmt.Errorf("user not found")
	}

	class, classExist := cly.store.GetClass(classId)
	if !classExist {
		return "", fmt.Errorf("class not found")
	}

	if class.ClassProviderUserName == userName {
		return "", fmt.Errorf("class provider cannot be class member")
	}

	bookingDate, err := utils.ParseTime(bookingDateStr)
	if err != nil {
		return "", fmt.Errorf("invalid booking date format")
	}

	classSessionId, err := cly.store.BookClass(userName, classId, bookingDate)
	if err != nil {
		return "", fmt.Errorf("error booking class: %w", err)
	}

	return classSessionId, nil
}

// GetVersion returns the current version of the Classly application.
func (cly *Classly) GetVersion() string {
	return "Classly-v0.1.0"
}
