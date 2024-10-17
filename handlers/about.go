package handlers

import (
	"net/http"
	"time"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
    p := &Page{
		Title:   "About Page",
		Message: "Hello, welcome to my simple Go website!",
		Year:    time.Now().Year(),
	}
    renderTemplate(w, "index.html", p)
}
