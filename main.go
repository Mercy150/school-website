package main

import (
	"html/template"
	"log"
	"net/http"
)

// Handler for the homepage
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Unable to load the homepage", http.StatusInternalServerError)
	}
}

// Handler for the application form
func applyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Process submitted form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to process form", http.StatusBadRequest)
			return
		}

		// Logging the submitted data
		log.Println("Student First Name:", r.FormValue("student_first_name"))
		log.Println("Parent/Guardian First Name:", r.FormValue("parent_first_name"))

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/apply.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Unable to load the application form", http.StatusInternalServerError)
	}
}

// Handler for Contact Us page
func contactHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/contact.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Unable to load the contact page", http.StatusInternalServerError)
	}
}

func main() {
	// Static file server
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/apply", applyHandler)
	http.HandleFunc("/contact", contactHandler)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
