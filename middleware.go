package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Middleware type
type Middleware func(http.Handler) http.Handler

// LogginMiddleware function
func LogginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)

		inMs := float64(duration) / float64(time.Millisecond)
		formatted := fmt.Sprintf("%.2f %s", inMs, "ms")

		log.Printf("- %s %s - %s", r.Method, r.URL.Path, formatted)
	})
}
