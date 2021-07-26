package todo

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/viper_config/internal/config"
)

type Config struct {
	*config.Config
}

func New(configuration *config.Config) *Config {
	return &Config{configuration}
}

func (config *Config) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{todoID}", config.GetATodo)
	router.Delete("/{todoID", config.DeleteTodo)
	router.Post("/", config.CreateTodo)
	router.Get("/", config.GetAllTodos)
	return router
}

func (config *Config) GetATodo(w http.ResponseWriter, r *http.Request)    {}
func (config *Config) DeleteTodo(w http.ResponseWriter, r *http.Request)  {}
func (config *Config) CreateTodo(w http.ResponseWriter, r *http.Request)  {}
func (config *Config) GetAllTodos(w http.ResponseWriter, r *http.Request) {}
