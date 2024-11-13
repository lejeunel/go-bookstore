package controllers

import (
	"github.com/google/uuid"
	g "go-bookstore/generic"
)

type AuthorInput struct {
	FirstName   string `json:"first_name" doc:"First name"`
	LastName    string `json:"last_name" doc:"Last name"`
	DateOfBirth string `json:"date_of_birth,omitempty" doc:"Date of birth"`
}

type AuthorMessage struct {
	Body AuthorInput
}

type AuthorOutputRecord struct {
	Id          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth string    `json:"date_of_birth,omitempty"`
}

type AuthorOutput struct {
	Body AuthorOutputRecord
}

type AuthorPaginatedOutputBody struct {
	Pagination *g.PaginationMeta    `json:"pagination"`
	Data       []AuthorOutputRecord `json:"data"`
}

type AuthorPaginatedOutput struct {
	Body AuthorPaginatedOutputBody
}

type GetOneAuthorInput struct {
	Id string `path:"id" doc:"ID to retrieve"`
}

type BookInput struct {
	Title string `json:"title" doc:"Title of book"`
}

type BookMessage struct {
	Body BookInput
}

type AuthorBookAssignInput struct {
	BookID   string `path:"book_id"`
	AuthorID string `path:"author_id"`
}

type BookOutputRecord struct {
	Id      uuid.UUID            `json:"id"`
	Title   string               `json:"title"`
	Authors []AuthorOutputRecord `json:"authors"`
}

type BookOutput struct {
	Body BookOutputRecord
}

type BookPaginatedOutputBody struct {
	Pagination *g.PaginationMeta  `json:"pagination"`
	Data       []BookOutputRecord `json:"data"`
}
type BookPaginatedOutput struct {
	Body BookPaginatedOutputBody
}

type GetOneBookInput struct {
	Id string `path:"id" doc:"ID to retrieve"`
}
