package main

import (
	"log"
	"math"
	"net/http"
	"runtime"
)

func main() {
	// Start CPU-intensive background tasks
	numCPU := runtime.NumCPU()
	log.Printf("Starting %d CPU-intensive goroutines", numCPU*2)

	for i := 0; i < numCPU*2; i++ {
		go func() {
			var result float64
			for {
				// Perform expensive calculations continuously
				for j := 0; j < 1000000; j++ {
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
