package app

import (
	"fmt"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/jmoiron/sqlx"
	c "go-bookstore/config"
	sql "go-bookstore/repositories/sql"
	routes "go-bookstore/routes"
	s "go-bookstore/services"
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

	bookRepo := sql.NewSQLBookRepo(a.DB)
	authorRepo := sql.NewSQLAuthorRepo(a.DB)

	bookService := s.BookService{BookRepo: bookRepo, AuthorRepo: authorRepo}
	authorService := s.AuthorService{AuthorRepo: authorRepo}

	a.Router = http.NewServeMux()
	api := humago.New(a.Router, huma.DefaultConfig("Book Store API", "1.0.0"))
	routes.AddRoutes(api, "/api/v1", bookService, authorService)
	a.Api = &api

}

func (a *App) Run(port int) {

	log.Printf("Starting server on port %d...\n", port)
	log.Printf("API docs: http://localhost:%d/docs\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), a.Router)
}
