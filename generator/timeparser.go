package generator

import (
	"time"
)

const TimeFormat = "2006-01-02T15:04:05"

var Location = loadTZ()

func ParseTime(s string) time.Time {
	t, err := time.ParseInLocation(TimeFormat, s, Location)
	if err != nil {
		panic(err)
	}
	return t
}

func loadTZ() *time.Location {
	loc, err := time.LoadLocation("Europe/Prague")
	if err != nil {
		panic(err)
	}
	return loc
}
