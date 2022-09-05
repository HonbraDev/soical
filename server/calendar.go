package server

import (
	"net/http"
	"time"

	"github.com/HonbraDev/soical/generator"

	ics "github.com/arran4/golang-ical"
)

func handleCalRequest(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	// auth
	username, password, ok := r.BasicAuth()
	if username == "" || password == "" || !ok {
		http.Error(w, "missing or improperly formatted credentials", http.StatusBadRequest)
		logRequest(r, "missing or improperly formatted credentials")
		return
	}

	// make
	cal, err := generator.GenerateCalendar(username, password)
	if err != nil {
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
