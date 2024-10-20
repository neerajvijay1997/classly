package utils

import (
	"fmt"
	"strings"
	"time"
)

// ParseTime converts a string to time.Time based on the provided layout (current layout: "2006-01-02")
func ParseTime(timeStr string) (time.Time, error) {
	parsedTime, err := time.Parse(DateFormat, timeStr)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}

func GetClassIdAndSessionDate(classSessionId string) (string, time.Time, error) {

	parts := strings.Split(classSessionId, "#")
	if len(parts) != 2 {
		return "", time.Time{}, fmt.Errorf("invalid sessionId format")

	}

	classId := parts[0]
	sessionDateStr := parts[1]

	sessionDate, err := time.Parse(DateFormat, sessionDateStr)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("invalid date format in sessionId: %v", err)
	}

	return classId, sessionDate, nil
}

func GenerateSessionId(classId string, bookingDate time.Time) string {
	formattedTime := bookingDate.Format(DateFormat)
	sessionId := fmt.Sprintf("%s#%s", classId, formattedTime)
	return sessionId
}
