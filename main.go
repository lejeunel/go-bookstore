package main

import (
	c "go-bookstore/config"
)

func main() {

	cfg := c.NewConfig()

	app := &c.App{}
	app.Initialize(cfg)

	app.Run(cfg.Port)
}
