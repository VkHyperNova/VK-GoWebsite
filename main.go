package main

import (
    "html/template"
    "log"
    "net/http"
)

// Page structure to hold data for the template
type Page struct {
    Title   string
    Message string
}

// Render the template with the provided data
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, err := template.ParseFiles("templates/" + tmpl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// Home handler to display the main page
func homeHandler(w http.ResponseWriter, r *http.Request) {
    p := &Page{Title: "Welcome", Message: "Hello, welcome to my simple Go website!"}
    renderTemplate(w, "index.html", p)
}

func main() {
    // Serve static files from the "static" directory
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // Set up routes
    http.HandleFunc("/", homeHandler)

    // Start the server
    log.Println("Starting server on :8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
