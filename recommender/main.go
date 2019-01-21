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
		&cobra.Command{
			Short: "Run an example query",
			Use:   "query",
			RunE:  cmd.Query,
		},
		&cobra.Command{
			Short: "Seed the database with some mock data",
			Use:   "seed",
			RunE:  cmd.Seed,
		},
		&cobra.Command{
			Short: "Train the recommendation model",
			Use:   "train",
			RunE:  cmd.Train,
		},
	)

	if err := root.Execute(); err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}
}
