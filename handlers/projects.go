package handlers

import (
	"net/http"
	"time"
)

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
    p := &Page{
		Title:   "Welcome Veikko Dev",
		Message: "Hello, welcome to my simple Go website!",
		Year:    time.Now().Year(),
	}
    renderTemplate(w, "index.html", p)
}