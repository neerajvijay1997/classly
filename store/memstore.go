package store

import (
	"classly/types"
	"fmt"
	"strings"
	"time"
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

func (ms *MemStore) updateUserForClassSession(userName string, bookingDate time.Time, classId string) {
	user := ms.users[userName]

	bookingDates, exist := user.BookedClasses[classId]

	if exist {
		user.BookedClasses[classId] = append(bookingDates, bookingDate)
	} else {
		user.BookedClasses[classId] = []time.Time{bookingDate}
	}

	ms.users[userName] = user
}

func (ms *MemStore) UpdateUserWithCreatedClass(userName string, classId string) {
	user := ms.users[userName]
	user.CreatedClassIds = append(user.CreatedClassIds, classId)
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

func (ms *MemStore) BookClass(userName string, classId string, bookingDate time.Time) (string, error) {
	classSessionId := generateSessionId(classId, bookingDate)

	session, ok := ms.classSessions[classSessionId]
	if ok {
		session = append(session, userName)
	} else {
		session = []string{userName}
	}

	ms.classSessions[classSessionId] = session

	ms.updateUserForClassSession(userName, bookingDate, classId)
	return classSessionId, nil
}

func (ms *MemStore) GetBookedClasses(userName string) ([]types.BookedClass, error) {
	var bookedClasses []types.BookedClass
	user, exist := ms.users[userName]
	if !exist {
		return bookedClasses, fmt.Errorf("user doesnt exist")
	}

	for classId, classSessions := range user.BookedClasses {
		class, exist := ms.GetClass(classId)
		if !exist {
			return bookedClasses, fmt.Errorf("class doesnot exist")

		}
		bookedClass := types.BookedClass{
			Class:    class,
			Sessions: classSessions,
		}
		bookedClasses = append(bookedClasses, bookedClass)
	}

	return bookedClasses, nil

}

func (ms *MemStore) GetClassesStatus(userName string) ([]types.ClassStatus, error) {
	var classesStatus []types.ClassStatus

	user, exist := ms.users[userName]
	if !exist {
		return classesStatus, fmt.Errorf("user doesnt exist")
	}

	for _, classId := range user.CreatedClassIds {
		class, exist := ms.GetClass(classId)
		if !exist {
			return classesStatus, fmt.Errorf("class doesnot exist")

		}

		classSessionsMap := make(types.ClassSessionsMap)
		currentDate := class.StartDate

		for {
			currentSessionId := generateSessionId(classId, currentDate)
			bookedUserNames, exist := ms.classSessions[currentSessionId]
			if exist {
				ms.populateClassSessionMap(classSessionsMap, currentSessionId, bookedUserNames)
			}

			if currentDate == class.EndDate {
				break
			} else {
				currentDate = currentDate.Add(24 * time.Hour)
			}
		}

		classStatus := types.ClassStatus{
			Class:    class,
			Sessions: classSessionsMap,
		}
		classesStatus = append(classesStatus, classStatus)

	}

	return classesStatus, nil
}

func (ms *MemStore) populateClassSessionMap(classSessionsMap types.ClassSessionsMap, classSessionId string, bookedUserNames []string) error {
	var users []types.User
	_, sessionDate, err := getClassIdAndSessionDate(classSessionId)
	if err != nil {
		return fmt.Errorf("error getting class id and session date")
	}

	for _, userNames := range bookedUserNames {
		user, _ := ms.GetUser(userNames)
		users = append(users, user)
	}

	classSessionsMap[sessionDate] = users
	return nil
}

func getClassIdAndSessionDate(classSessionId string) (string, time.Time, error) {

	parts := strings.Split(classSessionId, "#")
	if len(parts) != 2 {
		return "", time.Time{}, fmt.Errorf("invalid sessionId format")

	}

	classId := parts[0]
	sessionDateStr := parts[1]

	// TODO: Use constant
	sessionDate, err := time.Parse("2006-01-02", sessionDateStr)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("invalid date format in sessionId: %v", err)
	}

	return classId, sessionDate, nil
}

func generateSessionId(classId string, bookingDate time.Time) string {
	// TODO: Use constant
	formattedTime := bookingDate.Format("2006-01-02")
	sessionId := fmt.Sprintf("%s#%s", classId, formattedTime)
	return sessionId
}
