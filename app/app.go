package app

import (
	"context"
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
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	Api    *huma.API
	Router *http.ServeMux
	DB     *sqlx.DB
	Cfg    *c.Config
}

func NewApp(cfg *c.Config) *App {
	db := NewSQLiteConnection(cfg.Path)

	bookRepo := sql.NewSQLBookRepo(db)
	authorRepo := sql.NewSQLAuthorRepo(db)

	bookService := s.NewBookService(&bookRepo, &authorRepo, cfg.MaxPageSize, cfg.MaxPageSize)
	authorService := s.NewAuthorService(&authorRepo, &bookRepo, cfg.MaxPageSize, cfg.MaxPageSize)

	router := http.NewServeMux()
	api := humago.New(router, huma.DefaultConfig("Book Store API", "1.0.0"))
	routes.AddRoutes(api, "/api/v1", *bookService, *authorService)

	return &App{Api: &api, Router: router, DB: db, Cfg: cfg}

}

func (a *App) Run(port int) {
	server := &http.Server{Addr: fmt.Sprintf(":%d", port),
		Handler: a.Router}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	log.Printf("Starting server on port %d...\n", port)
	log.Printf("API docs URL: <root>:%d/docs\n", port)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	<-stopChan

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	log.Println("server stopping...")
	defer cancel()

	log.Fatal(server.Shutdown(ctx))
	log.Fatal(a.DB.Close())

}
