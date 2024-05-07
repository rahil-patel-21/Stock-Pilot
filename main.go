package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define the port to listen on
	port := ":8000" // You can change this to any port you prefer

	// Define your HTTP routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go!")
	})

	// Start the web server
	fmt.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
