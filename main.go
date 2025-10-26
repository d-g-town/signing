package main

import (
	"log"
	"net/http"
)

// authMiddleware enforces authentication on all routes
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for authorization header
		auth := r.Header.Get("Authorization")
		if auth != "Bearer secret-token" {
			w.Header().Set("WWW-Authenticate", `Bearer realm="restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			log.Printf("Blocked unauthorized request to: %s", r.URL.Path)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	// Root handler
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, authenticated user!"))
	})

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Wrap all routes with auth middleware
	handler := authMiddleware(mux)

	log.Println("Starting server on :8080")
	log.Println("All routes require Authorization: Bearer secret-token")

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
