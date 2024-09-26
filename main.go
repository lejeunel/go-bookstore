package main

import (
	"embed"
	"github.com/pressly/goose/v3"
	"go-bookstore/cmd"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {

	goose.SetBaseFS(embedMigrations)
	cmd.Execute()
}
