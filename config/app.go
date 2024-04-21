package config

import (
	"fmt"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"log"
	"net/http"
)

type App struct {
	Api    *huma.API
	Router *http.ServeMux
}

// Initialize app routing
func (a *App) Initialize(cfg *Config) {
	db := NewSQLiteConnection(cfg.Path)
	a.Router = http.NewServeMux()
	api := humago.New(a.Router, huma.DefaultConfig("Book Store API", "1.0.0"))
	AddRoutes(api, db)
	a.Api = &api

}

func (a *App) Run(port int) {

	log.Printf("Starting server on port %d...\n", port)
	log.Printf("API docs: http://localhost:%d/docs\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), a.Router)
}
