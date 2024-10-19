package classly

import (
	"classly/store"
	"classly/types"
	"fmt"
	"time"

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
	user := types.User{
		UserName: userName,
		Name:     name,
		Email:    email,
	}

	cly.store.SetUser(user)
	return userName
}

func (cly *Classly) GetUserInfo(userName string) (types.User, bool) {
	user, ok := cly.store.GetUser(userName)
	return user, ok
}

func (cly *Classly) CreateClass(userName, className string, startDateStr, endDateStr string, capacity uint32) (string, error) {
	// TODO: Check if the user exists

	startDate, err := ParseTime(startDateStr)
	if err != nil {
		return "", fmt.Errorf("invalid start date format")
	}

	endDate, err := ParseTime(endDateStr)
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

	class := types.Class{
		Id:                    classId,
		ClassName:             className,
		ClassProviderUserName: userName,
		Capacity:              capacity,
		StartDate:             startDate,
		EndDate:               endDate,
	}

	cly.store.SetClass(class)
	cly.store.AddOfferedClassToUser(userName, classId)

	return classId, nil
}

func (cly *Classly) GetAllClasses() []types.Class {
	return cly.store.GetAllClasses()
}

func (cly *Classly) GetOfferedClassesStatus(userName string) {

}

func (cly *Classly) GetEnrolledClassesStatus(userName string) {

}

func (cly *Classly) BookClass(userName string, classId string, bookingDateStr string) (string, error) {
	// TODO: Validate params
	// TODO: Class provider cannot be class member

	bookingDate, err := ParseTime(bookingDateStr)
	if err != nil {
		return "", fmt.Errorf("invalid booking date format")
	}

	classSessionId := generateSessionId(classId, bookingDate)
	cly.store.AddUserToClassSession(classSessionId, userName)

	return classSessionId, nil
}

func (cly *Classly) GetVersion() string {
	return "Classly-v0.1.0"
}

func generateSessionId(classId string, bookingDate time.Time) string {
	// TODO: Use constant
	formattedTime := bookingDate.Format("2006-01-02")
	sessionId := fmt.Sprintf("%s-%s", classId, formattedTime)
	return sessionId

}

// TODO: Move to utils package
// ParseTime converts a string to time.Time based on the provided layout (current layout: "2006-01-02")
func ParseTime(timeStr string) (time.Time, error) {
	// TODO: Use constant
	layout := "2006-01-02"
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}
