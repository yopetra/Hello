package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"example.com/mod/controllers"
	"example.com/mod/views"
	"github.com/go-chi/chi/v5"
)

// func executeTemplate(w http.ResponseWriter, filepath string) {
// t, err := views.Parse(filepath)
// if err != nil {
// log.Printf("parsing templat: %v, err")
// http.Error(w, "There was an error parsing the template.",
// http.StatusInternalServerError)
// return
// }
// t.Execute(w, nil)
// }

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// tplPath := filepath.Join("templates", "home.gohtml")
// executeTemplate(w, tplPath)
// }

// func contactHandler(w http.ResponseWriter, r *http.Request) {
// tplPath := filepath.Join("templates", "contact.gohtml")
// executeTemplate(w, tplPath)
// }

// func faqHandler(w http.ResponseWriter, r *http.Request) {
// executeTemplate(w, filepath.Join("templates", "faq.gohtml"))
// }

func main() {
	r := chi.NewRouter()
	tpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		panic(err)
	}

	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		panic(err)
	}

	r.Get("/faq", controllers.StaticHandler(tpl))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting server...")
	http.ListenAndServe(":3000", r)
}
