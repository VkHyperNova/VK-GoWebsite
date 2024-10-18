package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title   string
	Message string
	Year    int
}


func HomeHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Rendering Home Page")

	templates := template.Must(template.ParseFiles(
        "templates/layout.html",
        "templates/header.html",
        "templates/footer.html",
        "templates/home.html", // Specific projects page template
    ))

	data := Page{
		Title:   "Home",
		Message: "Welcome to the Home Page!",
		Year:    2024,
	}

	err := templates.ExecuteTemplate(w, "layout.html", data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Rendering About Page")

	templates := template.Must(template.ParseFiles(
        "templates/layout.html",
        "templates/header.html",
        "templates/footer.html",
        "templates/about.html", // Specific projects page template
    ))

	data := Page{
		Title:   "About",
		Message: "Welcome to the About Page!",
		Year:    2024,
	}
	err := templates.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Rendering Projects Page")

	templates := template.Must(template.ParseFiles(
        "templates/layout.html",
        "templates/header.html",
        "templates/footer.html",
        "templates/projects.html", // Specific projects page template
    ))

	data := Page{
		Title:   "Projects",
		Message: "Welcome to the Projects Page!",
		Year:    2024,
	}
	err := templates.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
