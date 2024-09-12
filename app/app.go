package app

import (
	"fmt"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/jmoiron/sqlx"
	c "go-bookstore/config"
	r "go-bookstore/repositories"
	routes "go-bookstore/routes"
	"log"
	"net/http"
)

type App struct {
	Api    *huma.API
	Router *http.ServeMux
	DB     *sqlx.DB
}

// Initialize app routing
func (a *App) Initialize(cfg *c.Config) {
	a.DB = NewSQLiteConnection(cfg.Path)
	paginator := &r.Paginator{MaxPageSize: cfg.MaxPageSize}
	a.Router = http.NewServeMux()
	api := humago.New(a.Router, huma.DefaultConfig("Book Store API", "1.0.0"))
	routes.AddRoutes(api, a.DB, paginator, "/api/v1")
	a.Api = &api

}

func (a *App) Run(port int) {

	log.Printf("Starting server on port %d...\n", port)
	log.Printf("API docs: http://localhost:%d/docs\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), a.Router)
}
