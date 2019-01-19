package cmd

import (
	"fmt"
	"popcorn/recommender/seeder"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Postgres Driver
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Seed inserts data to Postgres.
func Seed(cmd *cobra.Command, args []string) error {
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

	db, err := gorm.Open("postgres", addr)
	if err != nil {
		return err
	}

	logrus.Infof("connected to %s", addr)

	seeder.SetDatasetDir("./datasets/100k")
	if err := seeder.LoadMovies(); err != nil {
		return err
	}

	if err := seeder.LoadMetadata(); err != nil {
		return err
	}

	if err := seeder.LoadTags(); err != nil {
		return err
	}

	var count int
	for _, movie := range seeder.GetMovies() {
		if err := db.Create(movie).Error; err != nil {
			logrus.Error(err)
		} else {
			logrus.Infof("movie %d is inserted", movie.ID)
			count++
		}
	}

	logrus.Infof("inserted %d movies to database", count)

	return nil
}
