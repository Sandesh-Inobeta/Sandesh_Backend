package routers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"sandesh/controllers"
)

func Routes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Post("/getName", controllers.GetName)

	return router

}
