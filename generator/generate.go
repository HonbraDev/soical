package generator

import (
	"errors"
	"time"

	sogo "github.com/HonbraDev/sogo/client"
	ics "github.com/arran4/golang-ical"
)

func GenerateCalendar(username, password string) (*ics.Calendar, error) {
	// initialize client
	client := sogo.NewClient(username, password)
	auth, err := client.GetAuthStatus()
	if err != nil {
		return nil, err
	}
	if !auth {
		return nil, errors.New("authentication failed with no error")
	}

	// get user info
	userinfo, err := client.GetUzivatelInfo("")
	if err != nil {
		return nil, err
	}

	// fetch events
	events, err := client.GetRozvrhoveUdalostiRange(
		FormatDateYMD(time.Now().AddDate(0, 0, -7)),
		FormatDateYMD(time.Now().AddDate(0, 1, 0)),
	)
	if err != nil {
		return nil, err
	}

	// initialize calendar
	cal := ics.NewCalendar()
	cal.SetName("Rozvrh - " + userinfo.Jmeno)
	cal.SetTzid("Europe/Prague")
	cal.SetMethod("PUBLISH")
	cal.SetProductId("-//HonbraDev//SOiCal//EN")
	cal.SetCalscale("GREGORIAN")

	// add events
	for _, e := range events {
		ce := cal.AddEvent(e.UdalostID + "@" + e.CasOd + "@" + e.CasDo)
		ce.SetDtStampTime(time.Now())
		ce.SetSummary(e.Nazev)
		ce.SetStartAt(ParseTime(e.CasOd))
		ce.SetEndAt(ParseTime(e.CasDo))
		ce.SetLocation(FormatMistnostiUdalosti(e.MistnostiUdalosti))
	}

	return cal, nil
}
