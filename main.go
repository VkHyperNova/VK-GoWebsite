package main

import (
	"log"
	"net/http"
	"text/template"
)
// output.css and styles.css does not work
func main() {
	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
	// Set up routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/projects", projectsHandler)

	// Start the server
	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func renderTemplate(w http.ResponseWriter, templateName string) {
    tmpl, err := template.ParseFiles("templates/layout.html", "templates/header.html", "templates/footer.html", templateName)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.ExecuteTemplate(w, "layout", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "templates/home.html")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "templates/about.html")
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "templates/projects.html")
}
