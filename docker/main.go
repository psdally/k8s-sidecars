package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// register hello function to handle all requests
	mux := http.NewServeMux()

        http.Handle("/", http.FileServer(http.Dir("/www")))

	// use PORT environment variable, or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
