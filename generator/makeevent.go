package generator

import (
	"time"

	"github.com/HonbraDev/sogo/models"
	ics "github.com/arran4/golang-ical"
)

func MakeEvent(u *models.RozvrhovaUdalost, nameMap map[string]string) *ics.VEvent {
	e := ics.NewEvent(makeUniqueID(u))
	e.SetDtStampTime(time.Now())
	e.SetSummary(makeEventName(u, nameMap))
	e.SetStartAt(ParseTime(u.CasOd))
	e.SetEndAt(ParseTime(u.CasDo))
	e.SetLocation(FormatMistnostiUdalosti(u.MistnostiUdalosti))
	e.SetTimeTransparency(ics.TransparencyTransparent)
	for _, t := range u.UciteleUdalosti {
		e.AddAttendee(FormatUcitelUdalosti(&t))
	}
	for _, s := range u.SkupinyUdalosti {
		e.AddAttendee(FormatSkupinaUdalosti(&s))
	}
	return e
}

func makeUniqueID(u *models.RozvrhovaUdalost) string {
	return u.UdalostID + "@" + u.CasOd + "@" + u.CasDo
}

func makeEventName(u *models.RozvrhovaUdalost, nameMap map[string]string) string {
	if u.Predmet == nil {
		return u.Nazev
	}
	if name, ok := nameMap[u.Predmet.PredmetID]; ok {
		return name
	}
	if u.Predmet.Nazev != "" {
		return u.Predmet.Nazev
	}
	return u.Nazev
}
