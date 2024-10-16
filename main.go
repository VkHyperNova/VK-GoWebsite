package main

import (
	"html/template" // Package for parsing and executing HTML templates
	"log"           // Package for logging errors
	"net/http"      // Package for HTTP client and server
	"os"            // Package for file handling
	"regexp"        // Package for regular expressions
)

// Page represents a web page with a title and body content.

type Page struct {
	Title string // Title of the page
	Body  []byte // Body content of the page in bytes
}

/*
save writes the page's body to a text file named after the page's title.
Returns an error if the write operation fails.
*/
func (p *Page) save() error {
	filename := p.Title + ".txt"                // Construct the filename from the title
	return os.WriteFile(filename, p.Body, 0600) // Write the body to the file with permission 0600
}

// loadPage loads a page from a text file.
// Returns a pointer to the Page and an error if the file cannot be read.
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"         // Construct the filename from the title
	body, err := os.ReadFile(filename) // Read the file content
	if err != nil {
		return nil, err // Return an error if the file cannot be read
	}
	// Return a new Page instance with the loaded title and body
	return &Page{Title: title, Body: body}, nil
}

// viewHandler handles requests to view a page.
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title) // Attempt to load the page
	if err != nil {
		// Redirect to the edit page if the page does not exist
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p) // Render the view template with the loaded page
}

// editHandler handles requests to edit a page.
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title) // Attempt to load the page
	if err != nil {
		// If the page doesn't exist, create a new Page with the title
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p) // Render the edit template with the page
}

// saveHandler handles requests to save a page after editing.
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")                  // Get the body content from the form
	p := &Page{Title: title, Body: []byte(body)} // Create a Page instance with the title and body
	err := p.save()                              // Save the page to a file
	if err != nil {
		// If there is an error, respond with an internal server error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Redirect to the view page after saving
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// Load templates once at the start for better performance.
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

// renderTemplate renders an HTML template with the provided page data.
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl, p) // Execute the specified template with the page data
	if err != nil {
		// If there's an error in execution, respond with an internal server error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// validPath is a compiled regular expression to validate the URL path format.
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// makeHandler wraps a handler function to validate the URL path and extract the title.
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path) // Match the URL path against the valid pattern
		if m == nil {
			http.NotFound(w, r) // Respond with 404 if the path does not match
			return
		}
		fn(w, r, m[2]) // Call the handler function with the extracted title
	}
}

// main is the entry point of the application.
func main() {
	// Set up the routes for viewing, editing, and saving pages
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	// Start the server on port 8080 and log any fatal errors
	log.Fatal(http.ListenAndServe(":8080", nil))
}
