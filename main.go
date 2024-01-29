package main

import (
	"fmt"
	"net/http"

	"example.com/mod/controllers"
	"example.com/mod/models"
	"example.com/mod/templates"
	"example.com/mod/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS,
			"home.gohtml", "tailwind.gohtml",
		))))

	r.Get("/contact", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS,
			"contact.gohtml", "tailwind.gohtml",
		))))

	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS,
			"faq.gohtml", "tailwind.gohtml",
		))))

	// Setup a database connection
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Setup our model services
	userService := models.UserService{
		DB: db,
	}

	usersC := controllers.Users{
		UserService: &userService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS, "signup.gohtml", "tailwind.gohtml"))
	usersC.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS, "signin.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)
	r.Get("/signin", usersC.SignIn)

	r.Get("/welcome", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "welcome.gohtml"))))

	r.Get("/feedback", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS,
			"feedback.gohtml", "tailwind.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting server...")
	http.ListenAndServe(":3000", r)
}
