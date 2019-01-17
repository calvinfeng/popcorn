package cmd

import (
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Migrate runs migration on a PostgreSQL database.
func Migrate(cmd *cobra.Command, args []string) error {
	cfg := EnvConfig{}
	if err := env.Parse(&cfg); err != nil {
		return err
	}

	logrus.Info("performing migration on database")

	return nil
}
