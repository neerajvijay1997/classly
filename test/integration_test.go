package classly_test

import (
	"classly/classly"
	"classly/store"
	"classly/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestClasslyIntegration(t *testing.T) {
	// Initialize Classly instance
	memStore := store.NewMemStore()
	cly := classly.InitializeClassly(memStore)

	// Create multiple users
	user1Name, err := cly.CreateUser("Alice", "alice@example.com")
	require.NoError(t, err, "User creation should succeed")
	user2Name, err := cly.CreateUser("Bob", "bob@example.com")
	require.NoError(t, err, "User creation should succeed")

	// Verify users exist
	user1Info, user1Exists := cly.GetUserInfo(user1Name)
	require.True(t, user1Exists, "User1 should exist")
	require.Equal(t, user1Name, user1Info.UserName, "User1 name should match")

	user2Info, user2Exists := cly.GetUserInfo(user2Name)
	require.True(t, user2Exists, "User2 should exist")
	require.Equal(t, user2Name, user2Info.UserName, "User2 name should match")

	// User1 creates a class
	startDateStr := time.Now().Add(24 * time.Hour).Format(utils.DateFormat) // Class starts in 1 day
	endDateStr := time.Now().Add(96 * time.Hour).Format(utils.DateFormat)   // Class ends in 4 days
	classID, err := cly.CreateClass(user1Name, "Yoga Class", "Best yoga class available in the town", startDateStr, endDateStr, 10)
	require.NoError(t, err, "Class creation should succeed")
	require.NotEmpty(t, classID, "ClassID should not be empty")

	// User2 fetches all available classes
	classes, err := cly.GetAllClasses()
	require.NoError(t, err, "Fetching classes should succeed")
	require.Len(t, classes, 1, "There should be 1 class available")
	require.Equal(t, "Yoga Class", classes[0].ClassName, "Class name should match")

	// User2 books the class
	bookingDateStr := time.Now().Add(48 * time.Hour).Format(utils.DateFormat) // Booking time within class start and end
	bookingID, err := cly.BookClass(user2Name, classID, bookingDateStr)
	require.NoError(t, err, "Booking the class should succeed")
	require.NotEmpty(t, bookingID, "BookingID should not be empty")

	// User2 verifies the booking
	bookedClasses, err := cly.GetBookedClasses(user2Name)
	require.NoError(t, err, "Fetching booked classes should succeed")
	require.Len(t, bookedClasses, 1, "User2 should have 1 booked class")

	bookedClass := bookedClasses[0]
	require.Equal(t, classID, bookedClass.Id, "Booked class ID should match")
	require.Equal(t, "Yoga Class", bookedClass.ClassName, "Booked class name should match")
	require.Equal(t, user1Name, bookedClass.ClassProviderUserName, "Class provider should be User1")
	require.Len(t, bookedClass.Sessions, 1, "User2 should have booked 1 session")
	require.Equal(t, bookingDateStr, bookedClass.Sessions[0].Format(utils.DateFormat), "Booked session date should match")

	// User1 checks class status
	classStatuses, err := cly.GetClassesStatus(user1Name)
	require.NoError(t, err, "Fetching class status should succeed")
	require.Len(t, classStatuses, 1, "User1 should have 1 class status")

	classStatus := classStatuses[0]
	require.Equal(t, classID, classStatus.Id, "Class ID should match")
	require.Equal(t, "Yoga Class", classStatus.ClassName, "Class name should match")

	// Verify that User2 is in the session map for the class
	require.Len(t, classStatus.Sessions, 1, "There should be 1 session in the class")
	sessionUsers := classStatus.Sessions[bookedClass.Sessions[0]]
	require.Len(t, sessionUsers, 1, "There should be 1 user in the session")
	require.Equal(t, user2Name, sessionUsers[0].UserName, "User2 should be booked in the session")
	require.Equal(t, "Bob", sessionUsers[0].Name, "User2's name should match")
	require.Equal(t, "bob@example.com", sessionUsers[0].Email, "User2's email should match")
}
