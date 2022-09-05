package server

import (
	"log"
	"net/http"

	"github.com/HonbraDev/soical/generator"

	ics "github.com/arran4/golang-ical"
	"github.com/gorilla/mux"
)

func serveCalendar(w http.ResponseWriter, c *ics.Calendar) error {
	w.Header().Set("Content-Type", "text/calendar")
	w.Header().Set("charset", "utf-8")
	w.Header().Set("Content-Disposition", "inline")
	w.Header().Set("filename", "calendar.ics")
	return c.SerializeTo(w)
}

func handleCalRequest(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path, r.RemoteAddr)

	username, password, ok := r.BasicAuth()
	if username == "" || password == "" || !ok {
		http.Error(w, "missing or improperly formatted credentials", http.StatusBadRequest)
		return
	}

	cal, err := generator.GenerateCalendar(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := serveCalendar(w, cal); err != nil {
		log.Println(err)
	}
}

func NewServer(addr string) *http.Server {
	mux := mux.NewRouter().StrictSlash(true)
	mux.NewRoute().HandlerFunc(handleCalRequest).Methods("GET")
	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}
