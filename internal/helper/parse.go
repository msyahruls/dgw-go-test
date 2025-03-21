package helper

import (
	"time"
)

func ParseDate(dateStr string) time.Time {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, dateStr)
	return t
}
