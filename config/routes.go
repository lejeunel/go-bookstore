package config

import (
	"github.com/jmoiron/sqlx"
	"go-bookstore/controllers"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func AddRoutes(api huma.API, db *sqlx.DB) {
	bookHandler := controllers.NewBookHandler(db)

	// TODO add prefix to all paths here...

	// v1.Handle("/book/{id}", mw.JWTAuthMiddleware(http.HandlerFunc(bookHandler.DeleteBook))).Methods(http.MethodDelete)
	huma.Register(api, huma.Operation{
		OperationID: "get-book",
		Method:      http.MethodGet,
		Path:        "/books/{id}",
		Summary:     "Get a book by ID",
	}, bookHandler.GetOne)

	huma.Register(api, huma.Operation{
		OperationID:   "post-book",
		Method:        http.MethodPost,
		Path:          "/books",
		Summary:       "Add a new book",
		DefaultStatus: http.StatusCreated,
	}, bookHandler.Create)
}
