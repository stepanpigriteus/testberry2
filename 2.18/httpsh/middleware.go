package httpsh

import (
	"fmt"
	"net/http"
	"time"

	"grep/2.18/utils"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		utils.NewSlogger().Info(fmt.Sprintf("Started %s %s", r.Method, r.URL.Path))
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		utils.NewSlogger().Info(fmt.Sprintf("Completed %s %s in %v", r.Method, r.URL.Path, duration))
	})
}
