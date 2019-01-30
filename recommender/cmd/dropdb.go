package cmd

import (
	"github.com/golang-migrate/migrate"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// DropDB resets the database.
func DropDB(cmd *cobra.Command, args []string) error {
	if err := configureViper(); err != nil {
		return err
	}

	addr := pgAddress()
	migration, err := migrate.New(dir, addr)
	if err != nil {
		return err
	}

	err = migration.Drop()
	if err != nil {
		return err
	}

	logrus.Infof("%s has dropped all tables and data", addr)

	return nil
}
