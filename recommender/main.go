package main

import (
	"os"
	"popcorn/recommender/cmd"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
}

var root = &cobra.Command{
	Use:     "recommender",
	Example: "recommender serve",
}

func main() {
	root.AddCommand(
		&cobra.Command{
			Short: "Start gRPC server and have it serving client requests",
			Use:   "serve",
			RunE:  cmd.Serve,
		},
		&cobra.Command{
			Short: "Run migration on PostgreSQL",
			Use:   "migrate",
			RunE:  cmd.Migrate,
		},
	)

	if err := root.Execute(); err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}
}
