package classly_test

import (
	"classly/classly"
	"classly/store"
	"classly/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestClassCreation(t *testing.T) {
	// Initialize Classly instance
	memStore := store.NewMemStore()
	cly := classly.InitializeClassly(memStore)

	// Create a user
	user1Name, err := cly.CreateUser("Alice", "alice@example.com")
	require.NoError(t, err, "User creation should succeed")

	// Test cases
	tests := []struct {
		name           string
		userName       string
		className      string
		description    string
		startDateStr   string
		endDateStr     string
		capacity       uint32
		expectErr      bool
		expectedErrMsg string
	}{
		{
			name:         "Valid class creation",
			userName:     user1Name,
			className:    "Yoga Class",
			description:  "A relaxing yoga session",
			startDateStr: time.Now().Add(24 * time.Hour).Format(utils.DateFormat), // Valid future date
			endDateStr:   time.Now().Add(72 * time.Hour).Format(utils.DateFormat), // Valid end date
			capacity:     10,
			expectErr:    false,
		},
		{
			name:           "User does not exist",
			userName:       "NonExistentUser",
			className:      "Yoga Class",
			description:    "A relaxing yoga session",
			startDateStr:   time.Now().Add(24 * time.Hour).Format(utils.DateFormat),
			endDateStr:     time.Now().Add(74 * time.Hour).Format(utils.DateFormat),
			capacity:       10,
			expectErr:      true,
			expectedErrMsg: "user not found",
		},
		{
			name:           "Invalid start date format",
			userName:       user1Name,
			className:      "Yoga Class",
			description:    "A relaxing yoga session",
			startDateStr:   "invalid-date",
			endDateStr:     time.Now().Add(74 * time.Hour).Format(utils.DateFormat),
			capacity:       10,
			expectErr:      true,
			expectedErrMsg: "invalid start date format",
		},
		{
			name:           "Start date is in the past",
			userName:       user1Name,
			className:      "Yoga Class",
			description:    "A relaxing yoga session",
			startDateStr:   time.Now().Add(-1 * time.Hour).Format(utils.DateFormat), // Past date
			endDateStr:     time.Now().Add(1 * time.Hour).Format(utils.DateFormat),
			capacity:       10,
			expectErr:      true,
			expectedErrMsg: "start date must be in the future",
		},
		{
			name:           "End date is before start date",
			userName:       user1Name,
			className:      "Yoga Class",
			description:    "A relaxing yoga session",
			startDateStr:   time.Now().Add(24 * time.Hour).Format(utils.DateFormat),
			endDateStr:     time.Now().Add(1 * time.Hour).Format(utils.DateFormat), // Invalid end date
			capacity:       10,
			expectErr:      true,
			expectedErrMsg: "end date must be after start date",
		},
		{
			name:           "Invalid end date format",
			userName:       user1Name,
			className:      "Yoga Class",
			description:    "A relaxing yoga session",
			startDateStr:   time.Now().Add(24 * time.Hour).Format(utils.DateFormat),
			endDateStr:     "invalid-date",
			capacity:       10,
			expectErr:      true,
			expectedErrMsg: "invalid end date format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			classId, err := cly.CreateClass(tt.userName, tt.className, tt.description, tt.startDateStr, tt.endDateStr, tt.capacity)
			if tt.expectErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.expectedErrMsg)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, classId, "Class ID should not be empty")
			}
		})
	}
}

func TestBooking(t *testing.T) {
	// Initialize Classly instance
	memStore := store.NewMemStore()
	cly := classly.InitializeClassly(memStore)

	// Create multiple users
	user1Name, err := cly.CreateUser("Alice", "alice@example.com")
	require.NoError(t, err, "User creation should succeed")
	user2Name, err := cly.CreateUser("Bob", "bob@example.com")
	require.NoError(t, err, "User creation should succeed")

	// User creates a class
	startDateStr := time.Now().Add(24 * time.Hour).Format(utils.DateFormat) // Class starts in 1 day
	endDateStr := time.Now().Add(96 * time.Hour).Format(utils.DateFormat)   // Class ends in 4 days
	classID, err := cly.CreateClass(user1Name, "Yoga Class", "Best yoga class available in the town", startDateStr, endDateStr, 10)
	require.NoError(t, err, "Class creation should succeed")

	// Test cases
	tests := []struct {
		name           string
		userName       string
		classId        string
		bookingDateStr string
		expectErr      bool
		expectedErrMsg string
	}{
		{
			name:           "Valid booking",
			userName:       user2Name,
			classId:        classID,
			bookingDateStr: time.Now().Add(48 * time.Hour).Format(utils.DateFormat), // Booking time within class start and end
			expectErr:      false,
		},
		{
			name:           "User does not exist",
			userName:       "NonExistentUser",
			classId:        classID,
			bookingDateStr: time.Now().Add(48 * time.Hour).Format(utils.DateFormat),
			expectErr:      true,
			expectedErrMsg: "user not found",
		},
		{
			name:           "Class does not exist",
			userName:       user2Name,
			classId:        "NonExistentClassID",
			bookingDateStr: time.Now().Add(48 * time.Hour).Format(utils.DateFormat),
			expectErr:      true,
			expectedErrMsg: "class not found",
		},
		{
			name:           "Class provider cannot book their own class",
			userName:       user1Name,
			classId:        classID,
			bookingDateStr: time.Now().Add(48 * time.Hour).Format(utils.DateFormat),
			expectErr:      true,
			expectedErrMsg: "class provider cannot be class member",
		},
		{
			name:           "Booking date is outside class dates",
			userName:       user2Name,
			classId:        classID,
			bookingDateStr: time.Now().Add(120 * time.Hour).Format(utils.DateFormat), // Booking time after class ends
			expectErr:      true,
			expectedErrMsg: "booking date must be between",
		},
		{
			name:           "Invalid booking date format",
			userName:       user2Name,
			classId:        classID,
			bookingDateStr: "invalid-date",
			expectErr:      true,
			expectedErrMsg: "invalid booking date format",
		},
		{
			name:           "Booking exists for the provided date",
			userName:       user2Name,
			classId:        classID,
			bookingDateStr: time.Now().Add(48 * time.Hour).Format(utils.DateFormat), // Same booking date as before
			expectErr:      true,                                                    // Book the class first
			expectedErrMsg: "booking exist for provided date",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bookingID, err := cly.BookClass(tt.userName, tt.classId, tt.bookingDateStr)
			if tt.expectErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.expectedErrMsg)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, bookingID, "Booking ID should not be empty")
			}
		})
	}
}
