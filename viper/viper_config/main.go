package main

import (
	"log"

	"github.com/go-chi/chi"
	"github.com/viper_config/internal/config"
)

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/todo", todo.New(configuration).Routes())
	})

	return router
}

func main() {
	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error", err)
	}
	router := Routes(configuration)
}
