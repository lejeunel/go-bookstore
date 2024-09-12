package cmd

import (
	"github.com/spf13/cobra"
	a "go-bookstore/app"
	c "go-bookstore/config"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run server",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	cfg := c.NewConfig()
	app := &a.App{}
	app.Initialize(cfg)

	app.Run(cfg.Port)
}
