package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/alochym01/web-w-golang/project/driver"
	"github.com/alochym01/web-w-golang/project/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	connection, err := driver.ConnectSQL("127.0.0.1", 3306, "alochym", "Alochym@123", "alochym")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	postHandler := handler.NewPostHandler(connection)

	r.Route("/", func(rt chi.Router) {
		rt.Mount("/posts", postRouter(postHandler))
	})

	fmt.Println("Server listen at :8080")
	http.ListenAndServe(":8080", r)

}

// A completely separate router for posts routes
func postRouter(pHandler *handler.PostController) http.Handler {
	r := chi.NewRouter()
	r.Get("/", pHandler.Fetch)
	r.Get("/{id:[0-9]+}", pHandler.GetByID)
	r.Post("/", pHandler.Create)
	r.Put("/{id:[0-9]+}", pHandler.Update)
	r.Delete("/{id:[0-9]+}", pHandler.Delete)

	return r
}
