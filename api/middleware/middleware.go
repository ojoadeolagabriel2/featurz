package middleware

import "net/http"

func CommonMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Application", "featurz")
		handler.ServeHTTP(w, r)
	})
}