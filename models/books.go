package models

import (
	"github.com/google/uuid"
)

type Book struct {
	Id      uuid.UUID `db:"id"`
	Title   string    `db:"title"`
	Authors []Author
}
