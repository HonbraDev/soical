package server

import (
	"net/http"
	"net/url"

	"github.com/HonbraDev/soical/shared"
)

func logRequestHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest(r, "Started")
		handler.ServeHTTP(w, r)
	})
}

func logRequest(r *http.Request, args ...any) {
	u := url.URL(*r.URL)
	q := u.Query()
	if q.Get("password") != "" {
		q.Set("password", "REDACTED")
	}
	u.RawQuery = q.Encode()
	args = append([]any{r.RemoteAddr, r.Method, u.String()}, args...)
	shared.L.Println(args...)
}
