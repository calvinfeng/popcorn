package cmd

import (
	"fmt"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres" // Driver
	_ "github.com/golang-migrate/migrate/source/file"       // Driver
	_ "github.com/lib/pq"                                   // Driver
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func pgAddress() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.hostname"),
		viper.GetInt("database.port"),
		viper.GetString("database.name"),
		viper.GetString("database.ssl_mode"),
	)
}

const dir = "file://./migrations/"

// Migrate runs migration on a PostgreSQL database.
func Migrate(cmd *cobra.Command, args []string) error {
	if err := configureViper(); err != nil {
		return err
	}

	addr := pgAddress()
	migration, err := migrate.New(dir, pgAddress())
	if err != nil {
		return err
	}

	if err := migration.Up(); err != nil && err == migrate.ErrNoChange {
		logrus.Warn(err)
	} else if err != nil {
		return err
	}

	logrus.Infof("%s has migrated to latest version", addr)

	return nil
}
