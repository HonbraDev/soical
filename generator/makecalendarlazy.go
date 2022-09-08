package generator

import (
	"time"

	sogo "github.com/HonbraDev/sogo/client"
	ics "github.com/arran4/golang-ical"
)

func MakeCalendarLazy(username, password string) (*ics.Calendar, error) {
	client := sogo.NewClient(username, password)

	events, err := client.GetRozvrhoveUdalostiRange(
		FormatDateYMD(time.Now().AddDate(0, 0, -7)),
		FormatDateYMD(time.Now().AddDate(0, 1, 0)),
	)
	if err != nil {
		return nil, err
	}

	return MakeCalendar(events, nil)
}
