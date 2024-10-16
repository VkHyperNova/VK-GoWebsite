package handlers

import (
    "net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    p := &Page{Title: "Welcome Veikko Dev", Message: "Hello, welcome to my simple Go website!"}
    renderTemplate(w, "index.html", p)
}
