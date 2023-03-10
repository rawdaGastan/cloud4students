// Package middlewares for middleware between api and backend
package middlewares

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

// LoggingMW logs all information of every request
func LoggingMW(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msgf("%v: %v\n%v", r.Method, r.RequestURI, time.Now().Format(time.RFC850))
		h.ServeHTTP(w, r)
	})
}
