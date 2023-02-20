package main

import (
	"API/API"
	"API/backend"
	"fmt"
	"net/http"
)

func main() {
	// Serve static files from the "front-end" directory
	fs := http.FileServer(http.Dir("front-end"))
	http.Handle("/front-end/", http.StripPrefix("/front-end/", fs))

	// Set up API endpoints
	http.HandleFunc("/", backend.HandleTest)
	http.HandleFunc("/dates", API.HandleDates)
	http.HandleFunc("/artist", API.HandleArtists)
	http.HandleFunc("/data", API.HandleData)
	http.HandleFunc("/location", API.HandleLocations)
	http.HandleFunc("/info", API.HandleInfo)
	http.HandleFunc("/relation", API.HandleRelations)

	http.HandleFunc("/nwa", backend.HandleInfoNWA)
	http.HandleFunc("/shakur", backend.HandleInfoShakur)
	http.HandleFunc("/notorious", backend.HandleInfoNotorious)
	http.HandleFunc("/lamar", backend.HandleInfoLamar)
	http.HandleFunc("/warren", backend.HandleInfoWarren)
	http.HandleFunc("/daft", backend.HandleInfoDaftPunk)

	// Start server
	fmt.Println("Server started on port :8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}
