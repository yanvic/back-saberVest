package routes

import (
	"back-sabervest/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type Response struct {
	Msg  string
	Code int
}

func CreateRouter() *chi.Mux {

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CRSF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Route("/v1", func(router chi.Router) {

		router.Get("/college", handlers.GetUniversity)
		//router.Get("/{college}/{matter}", handlers.GetCollege)
		router.Get("/questions_all", handlers.GetQuestions)
		router.Get("/questions", handlers.GetQuestionsParams)

	})

	return router

}
