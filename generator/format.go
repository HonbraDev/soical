package generator

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
