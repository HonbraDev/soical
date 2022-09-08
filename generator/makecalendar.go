package generator

import (
	"github.com/HonbraDev/sogo/models"
	ics "github.com/arran4/golang-ical"
)

func MakeCalendar(events []models.RozvrhovaUdalost, nameMap map[string]string) (*ics.Calendar, error) {
	cal := ics.NewCalendar()
	cal.SetTzid("Europe/Prague")
	cal.SetMethod("PUBLISH")
	cal.SetProductId("-//HonbraDev//SOiCal//EN")
	cal.SetCalscale("GREGORIAN")
	for _, e := range events {
		cal.AddVEvent(MakeEvent(&e, nameMap))
	}
	return cal, nil
}
