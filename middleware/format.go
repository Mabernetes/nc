package middlewares

import (
	"encoding/json"
	"net/http"
)

type JSONWriter struct {
	http.ResponseWriter
	status int
	data   interface{}
}

func (rw *JSONWriter) WriteHeader(status int) {
	if rw.status != 0 {
		rw.status = status
		rw.ResponseWriter.WriteHeader(status)
	}
}

func (rw *JSONWriter) Write(b []byte) (int, error) {
	return 0, json.Unmarshal(b, &rw.data)
}

func JSONResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wr := &JSONWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(wr, r)

		if wr.data != nil {
			var response map[string]interface{}
			switch {
			case wr.status >= 200 && wr.status < 400:
				response = map[string]interface{}{"data": wr.data}
			case wr.status >= 400 && wr.status < 600:
				response = map[string]interface{}{"message": wr.data}
			default:
				response = map[string]interface{}{"data": wr.data}
			}
			response["status"] = wr.status

			wr.Header().Set("Content-Type", "application/json")
			w.WriteHeader(wr.status)
			if err := json.NewEncoder(w).Encode(response); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	})
}
