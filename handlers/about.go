package handlers

import "net/http"

func AboutHandler(w http.ResponseWriter, r *http.Request) {
    p := &Page{Title: "About", Message: "You are on about page!"}
    renderTemplate(w, "index.html", p)
}
