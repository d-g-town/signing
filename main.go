package main

import (
	"log"
	"math"
	"net/http"
	"os"
	"runtime"
	"strconv"
)

func main() {
	// Get configuration from environment variables
	numGoroutines := 1000
	if env := os.Getenv("NUM_GOROUTINES"); env != "" {
		if n, err := strconv.Atoi(env); err == nil {
			numGoroutines = n
		}
	}

	workload := 1000000
	if env := os.Getenv("WORKLOAD_SIZE"); env != "" {
		if n, err := strconv.Atoi(env); err == nil {
			workload = n
		}
	}

	// Start CPU-intensive background tasks
	log.Printf("Starting %d CPU-intensive goroutines with workload size %d", numGoroutines, workload)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			var result float64
			for {
				// Perform expensive calculations continuously
				for j := 0; j < workload; j++ {
					result += math.Sqrt(float64(j)) * math.Sin(float64(j))
				}
			}
		}()
	}

	mux := http.NewServeMux()

	// Catch-all handler for all routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Write([]byte("Hello, user!"))
			return
		}
		if r.URL.Path == "/health" {
			w.Write([]byte("OK"))
			return
		}
		// Return 404 for other paths
		http.NotFound(w, r)
	})

	log.Println("Starting server on :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
