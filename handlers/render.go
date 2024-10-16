package handlers

import (
    "html/template"
    "net/http"
)

// Page structure for templates
type Page struct {
    Title   string
    Message string
}

var templates = template.Must(template.ParseFiles(
    "templates/header.html",
    "templates/footer.html",
    "templates/index.html",
))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}