package generator

import "time"

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

func FormatDateYMD(date time.Time) string {
	return date.Format("2006-01-02")
}
