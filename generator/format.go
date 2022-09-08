package generator

import (
	"time"

	"github.com/HonbraDev/sogo/models"
)

func FormatMistnostiUdalosti(mistnostiUdalosti []models.MistnostUdalosti) string {
	loc := ""
	for i, m := range mistnostiUdalosti {
		if i > 0 {
			loc += ", "
		}
		loc += m.Nazev
	}
	return loc
}

func FormatUcitelUdalosti(u *models.UcitelUdalosti) string {
	return u.Jmeno + " " + u.Prijmeni
}

func FormatSkupinaUdalosti(s *models.SkupinaUdalosti) string {
	if s.SkupinaNazev != s.TridaNazev {
		return s.SkupinaNazev + " (" + s.TridaNazev + ")"
	} else {
		return s.SkupinaNazev
	}
}

func FormatDateYMD(date time.Time) string {
	return date.Format("2006-01-02")
}
