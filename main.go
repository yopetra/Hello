package main

import (
	"fmt"
	"net/http"

	"example.com/mod/controllers"
	"example.com/mod/templates"
	"example.com/mod/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "layout-page.gohtml", "home-page.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "layout-page.gohtml", "contact-page.gohtml"))))

	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml"))))

	r.Get("/welcome", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "welcome.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting server...")
	http.ListenAndServe(":3000", r)
}
