package models

import (
	"github.com/google/uuid"
)

type Book struct {
	Id      uuid.UUID `db:"id"`
	Title   string    `db:"title"`
	Authors []Author
}

type BookInputBody struct {
	Title string `json:"title" doc:"Title of book"`
}

type BookInput struct {
	Body BookInputBody
}

type AuthorBookAssign struct {
	BookID   string
	AuthorID string
}

type BookOutputRecord struct {
	Id      uuid.UUID            `json:"id"`
	Title   string               `json:"title"`
	Authors []AuthorOutputRecord `json:"authors"`
}

type BookOutput struct {
	Body BookOutputRecord
}

type BookPaginatedOutput struct {
	Body struct {
		Pagination *Pagination        `json:"pagination"`
		Data       []BookOutputRecord `json:"data"`
	}
}

type GetOneBookInput struct {
	Id string `path:"id" doc:"ID to retrieve"`
}
