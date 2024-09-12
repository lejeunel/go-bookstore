package cmd

import (
	migrationCmd "github.com/pivaldi/db-migration/cmds"
	migrationCfg "github.com/pivaldi/db-migration/config"
	"github.com/spf13/cobra"
	a "go-bookstore/app"
	c "go-bookstore/config"
)

var migrateCmd = &cobra.Command{
	Use:              "migration",
	Short:            "Database migration",
	TraverseChildren: true,
}

func init() {
	myCfg := c.NewConfig()

	a.NewSQLiteConnection(myCfg.Path)

	dbCfg := migrationCfg.ConfigT{DBConnection: migrationCfg.DBConnectionT{DSN: myCfg.Path,
		Driver: "sqlite3"}, DBMigration: migrationCfg.DBMigrationT{Dir: "migrations"}}

	migrationCmd.SetConfig(dbCfg)
	migrateCmd.AddCommand(migrationCmd.Root.Commands()...)
}
