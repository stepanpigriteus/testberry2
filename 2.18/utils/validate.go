package utils

import (
	"grep/2.18/domain"
	"strconv"
	"time"
)

func ValidStruct(resp domain.Event) bool {
	if resp == (domain.Event{}) || resp.Date.IsZero() || resp.UserID == 0 || resp.Event == "" {
		return false
	}
	return true
}

func ParseUserID(userID string) (int, error) {
	return strconv.Atoi(userID)
}

func ParseDate(date string) (time.Time, error) {
	const layout = "2006-01-02"
	return time.Parse(layout, date)
}
