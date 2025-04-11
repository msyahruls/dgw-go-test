package helper

import (
	"errors"
	"strconv"
	"time"
)

func ParseDate(dateStr string) time.Time {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, dateStr)
	return t
}

func ParseUintParam(param string) (uint, error) {
	id64, err := strconv.ParseUint(param, 10, 32)
	if err != nil {
		return 0, errors.New("ID must be a valid number")
	}
	return uint(id64), nil
}
