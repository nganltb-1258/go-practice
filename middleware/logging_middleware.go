package middleware

import (
    "log"
    "net/http"
    "time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("___________________________________________")
        start := time.Now() // Record start time
        log.Printf("Started %s %s", r.Method, r.URL.Path)

        next.ServeHTTP(w, r) // Call the next handler

        duration := time.Since(start) // Calculate duration
        log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, duration)
    })
}
