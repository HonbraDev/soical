package generator

import (
	"time"

	"github.com/HonbraDev/sogo/models"
	ics "github.com/arran4/golang-ical"
)

func GenerateEvent(u *models.RozvrhovaUdalost) *ics.VEvent {
	e := ics.NewEvent(makeUniqueID(u))
	e.SetDtStampTime(time.Now())
	if u.Predmet.Nazev != "" {
		e.SetSummary(u.Predmet.Nazev)
	} else {
		e.SetSummary(u.Nazev)
	}
	e.SetStartAt(ParseTime(u.CasOd))
	e.SetEndAt(ParseTime(u.CasDo))
	e.SetLocation(FormatMistnostiUdalosti(u.MistnostiUdalosti))
	e.SetTimeTransparency(ics.TransparencyTransparent)
	for _, t := range u.UciteleUdalosti {
		e.AddAttendee(FormatUcitelUdalosti(t))
	}
	for _, s := range u.SkupinyUdalosti {
		e.AddAttendee(FormatSkupinaUdalosti(s))
	}
	return e
}

func makeUniqueID(u *models.RozvrhovaUdalost) string {
	return u.UdalostID + "@" + u.CasOd + "@" + u.CasDo
}
