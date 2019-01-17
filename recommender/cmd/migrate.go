package cmd

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres" // Driver
	_ "github.com/golang-migrate/migrate/source/file"       // Driver
	_ "github.com/lib/pq"                                   // Driver
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	dir   = "file://./migrations/"
	up    = "up"
	reset = "reset"
)

const usage = `
Commands:
	up     Migrate the DB to the most recent version available
	reset  Resets the database
Usage:
	userauth migrate <command>
`

// Migrate runs migration on a PostgreSQL database.
func Migrate(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		fmt.Println(usage)
		return errors.New("no commands provided")
	}

	if err := configureViper(); err != nil {
		return err
	}

	addr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.hostname"),
		viper.GetInt("database.port"),
		viper.GetString("database.name"),
		viper.GetString("database.ssl_mode"),
	)

	logrus.Infof("performing migration on database %s", addr)

	migration, err := migrate.New(dir, addr)
	if err != nil {
		return err
	}

	switch args[0] {
	case up:
		return migration.Up()
	case reset:
		return migration.Drop()
	default:
		return fmt.Errorf("%s is not a valid command", args[0])
	}
}
