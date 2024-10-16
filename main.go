package main

import (
    "log"
    "net/http"
	"VK-GoWebsite/handlers"
)

// Page structure to hold data for the template
type Page struct {
    Title   string
    Message string
}

func main() {
    // Serve static files from the "static" directory
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // Set up routes
    http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

    // Start the server
    log.Println("Starting server on :8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
