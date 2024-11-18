package middlewares

import (
	"log"
	"net/http"
)

type LoggerWriter struct {
	http.ResponseWriter
	status int
}

func (rw *LoggerWriter) WriteHeader(status int) {
	if rw.status != 0 {
		rw.status = status
		rw.ResponseWriter.WriteHeader(status)
	}
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wr := &LoggerWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(wr, r)
		log.Printf("%d %s %d %s - %s - %s\n", wr.status, r.Method, r.ContentLength, r.URL, r.RemoteAddr, r.UserAgent())
	})
}
