package store

import (
	"classly/utils"
	"fmt"
	"time"
)

type MemStore struct {
	users         map[string]utils.User
	classes       map[string]utils.Class
	classSessions map[string][]string
}

func NewMemStore() *MemStore {
	return &MemStore{
		users:         make(map[string]utils.User),
		classes:       make(map[string]utils.Class),
		classSessions: make(map[string][]string),
	}
}

func (ms *MemStore) SetUser(user utils.User) error {
	ms.users[user.UserName] = user
	return nil
}

func (ms *MemStore) updateUserForBookedClassSession(userName string, bookingDate time.Time, classId string) {
	user := ms.users[userName]

	bookingDates, exist := user.BookedClasses[classId]

	if exist {
		user.BookedClasses[classId] = append(bookingDates, bookingDate)
	} else {
		user.BookedClasses[classId] = []time.Time{bookingDate}
	}

	ms.users[userName] = user
}

func (ms *MemStore) UpdateUserWithCreatedClass(userName string, classId string) error {
	user := ms.users[userName]
	user.CreatedClassIds = append(user.CreatedClassIds, classId)
	ms.users[userName] = user
	return nil
}

func (ms *MemStore) GetUser(userName string) (utils.User, bool) {
	user, ok := ms.users[userName]
	return user, ok
}

func (ms *MemStore) SetClass(class utils.Class) error {
	ms.classes[class.Id] = class
	return nil
}

func (ms *MemStore) GetClass(classId string) (utils.Class, bool) {
	class, ok := ms.classes[classId]
	return class, ok
}

func (ms *MemStore) GetAllClasses() ([]utils.Class, error) {
	allClasses := make([]utils.Class, 0, len(ms.classes))
	for _, class := range ms.classes {
		allClasses = append(allClasses, class)
	}
	return allClasses, nil
}

func (ms *MemStore) BookClass(userName string, classId string, bookingDate time.Time) (string, error) {
	classSessionId := utils.GenerateSessionId(classId, bookingDate)

	session, ok := ms.classSessions[classSessionId]
	if ok {
		session = append(session, userName)
	} else {
		session = []string{userName}
	}

	ms.classSessions[classSessionId] = session

	ms.updateUserForBookedClassSession(userName, bookingDate, classId)
	return classSessionId, nil
}

func (ms *MemStore) GetBookedClasses(userName string) ([]utils.BookedClass, error) {
	var bookedClasses []utils.BookedClass

	user, exist := ms.GetUser(userName)
	if !exist {
		return bookedClasses, fmt.Errorf("user doesnot exist")
	}

	for classId, classSessions := range user.BookedClasses {
		class, exist := ms.GetClass(classId)
		if !exist {
			return bookedClasses, fmt.Errorf("class doesnot exist")

		}
		bookedClass := utils.BookedClass{
			Class:    class,
			Sessions: classSessions,
		}
		bookedClasses = append(bookedClasses, bookedClass)
	}

	return bookedClasses, nil

}

func (ms *MemStore) GetClassesStatus(userName string) ([]utils.ClassStatus, error) {
	var classesStatus []utils.ClassStatus

	user, exist := ms.GetUser(userName)
	if !exist {
		return classesStatus, fmt.Errorf("user doesnot exist")
	}

	for _, classId := range user.CreatedClassIds {
		class, exist := ms.GetClass(classId)
		if !exist {
			return classesStatus, fmt.Errorf("class doesnot exist")

		}

		classSessionsMap := make(utils.ClassSessionsMap)
		currentDate := class.StartDate

		for {
			currentSessionId := utils.GenerateSessionId(classId, currentDate)
			bookedUserNames, exist := ms.classSessions[currentSessionId]
			if exist {
				err := ms.populateClassSessionMap(classSessionsMap, currentSessionId, bookedUserNames)
				if err != nil {
					return classesStatus, fmt.Errorf("error populating class session map: %w", err)
				}
			}

			if currentDate == class.EndDate {
				break
			} else {
				currentDate = currentDate.Add(24 * time.Hour)
			}
		}

		classStatus := utils.ClassStatus{
			Class:    class,
			Sessions: classSessionsMap,
		}
		classesStatus = append(classesStatus, classStatus)

	}

	return classesStatus, nil
}

func (ms *MemStore) populateClassSessionMap(classSessionsMap utils.ClassSessionsMap, classSessionId string, bookedUserNames []string) error {
	var users []utils.User
	_, sessionDate, err := utils.GetClassIdAndSessionDate(classSessionId)
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
