package generator

import (
	"time"
)

// pain.
func FormatMistnostiUdalosti(mistnostiUdalosti []struct {
	MistnostID     string `json:"MISTNOST_ID"`
	Nazev          string `json:"NAZEV"`
	Popis          string `json:"POPIS"`
	PriznakAbsence bool   `json:"PRIZNAK_ABSENCE"`
},
) string {
	loc := ""
	for i, m := range mistnostiUdalosti {
		if i > 0 {
			loc += ", "
		}
		loc += m.Nazev
	}
	return loc
}

func FormatUcitelUdalosti(u struct {
	UcitelID       string `json:"UCITEL_ID"`
	Prijmeni       string `json:"PRIJMENI"`
	Jmeno          string `json:"JMENO"`
	Zkratka        string `json:"ZKRATKA"`
	PriznakAbsence bool   `json:"PRIZNAK_ABSENCE"`
},
) string {
	return u.Jmeno + " " + u.Prijmeni
}

func FormatSkupinaUdalosti(s struct {
	SkupinaID          string `json:"SKUPINA_ID"`
	SkupinaNazev       string `json:"SKUPINA_NAZEV"`
	PriznakDruhSkupiny string `json:"PRIZNAK_DRUH_SKUPINY"`
	TridaID            string `json:"TRIDA_ID"`
	TridaNazev         string `json:"TRIDA_NAZEV"`
	PriznakAbsence     bool   `json:"PRIZNAK_ABSENCE"`
},
) string {
	if s.SkupinaNazev != s.TridaNazev {
		return s.SkupinaNazev + " (" + s.TridaNazev + ")"
	} else {
		return s.SkupinaNazev
	}
}

func FormatDateYMD(date time.Time) string {
	return date.Format("2006-01-02")
}
