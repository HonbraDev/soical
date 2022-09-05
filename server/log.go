package server

import (
	"log"
	"net/http"
)

func logRequestHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest(r, "Started")
		handler.ServeHTTP(w, r)
	})
}

func logRequest(r *http.Request, args ...any) {
	args = append([]any{r.RemoteAddr, r.Method, r.URL}, args...)
	log.Println(args...)
}
