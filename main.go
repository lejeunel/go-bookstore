package main

import (
	a "go-bookstore/app"
	c "go-bookstore/config"
)

func main() {
	cfg := c.NewConfig()
	app := &a.App{}
	app.Initialize(cfg)

	app.Run(cfg.Port)
}
