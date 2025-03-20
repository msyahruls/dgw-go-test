package handler

import (
	"time"
)

func parseDate(dateStr string) time.Time {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, dateStr)
	return t
}
