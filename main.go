package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// HealthResponse is a struct to define the structure of the JSON response.
type HealthResponse struct {
	Status string `json:"status"`
}

// healthCheckHandler writes the JSON response for the health check.
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// 2. Create the response data
	response := HealthResponse{Status: "HEALTHY"}

	// 3. Encode the struct into JSON and write it to the ResponseWriter
	// json.NewEncoder(w).Encode(response) is a common and efficient way
	// to write JSON directly to the response stream.
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}

	// For a simple health check, the default status code is 200 OK.
	// You don't need to explicitly call w.WriteHeader(http.StatusOK)
	// unless you are setting a different status code.
}

func main() {
	// Register the handler function for the root path
	http.HandleFunc("/", healthCheckHandler)

	// Define the port to listen on.
	const port = ":80"

	log.Printf("Starting server on port %s", port)

	// Start the HTTP server.
	// Note: Binding to port 80 on Linux/macOS usually requires root privileges (sudo).
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
