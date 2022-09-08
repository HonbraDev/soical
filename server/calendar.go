package server

import (
	"fmt"
	"net/http"
	"time"

	sogo "github.com/HonbraDev/sogo/client"
	"github.com/HonbraDev/soical/generator"

	ics "github.com/arran4/golang-ical"
)

func HandleCalRequest(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	// auth
	username, password, ok := r.BasicAuth()
	if username == "" || password == "" || !ok {
		username = r.URL.Query().Get("username")
		password = r.URL.Query().Get("password")
		if username == "" || password == "" {
			http.Error(w, "missing or improperly formatted credentials", http.StatusBadRequest)
			logRequest(r, "missing or improperly formatted credentials")
			return
		}
	}

	// name map
	var nameMap map[string]string
	if nameMapUrl := r.URL.Query().Get("nameMapUrl"); nameMapUrl != "" {
		var err error
		nameMap, err = generator.GetNameMap(nameMapUrl)
		if err != nil {
			err = fmt.Errorf("failed to get name map: %w", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logRequest(r, "Error:", err.Error())
			return
		}
	}

	// events
	client := sogo.NewClient(username, password)
	events, err := client.GetRozvrhoveUdalostiRange(
		generator.FormatDateYMD(time.Now().AddDate(0, 0, -7)),
		generator.FormatDateYMD(time.Now().AddDate(0, 1, 0)),
	)
	if err != nil {
		err = fmt.Errorf("failed to get events: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logRequest(r, "Error:", err.Error())
		return
	}

	// make
	cal, err := generator.MakeCalendar(events, nameMap)
	if err != nil {
		err = fmt.Errorf("failed to make calendar: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logRequest(r, "Error:", err)
		return
	}

	// send
	if err := serveCalendar(w, cal); err != nil {
		logRequest(r, "Error:", err)
		return
	}

	logRequest(r, "Done, took", time.Since(startTime))
}

func serveCalendar(w http.ResponseWriter, c *ics.Calendar) error {
	w.Header().Set("Content-Type", "text/calendar")
	w.Header().Set("charset", "utf-8")
	w.Header().Set("Content-Disposition", "inline")
	w.Header().Set("filename", "calendar.ics")
	return c.SerializeTo(w)
}
