package server

import (
	"net/http"

	"github.com/HonbraDev/soical/shared"
)

func logRequestHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest(r, "Started")
		handler.ServeHTTP(w, r)
	})
}

func logRequest(r *http.Request, args ...any) {
	args = append([]any{r.RemoteAddr, r.Method, r.URL}, args...)
	shared.L.Println(args...)
}
