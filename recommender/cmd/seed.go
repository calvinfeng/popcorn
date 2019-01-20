package cmd

import (
	"fmt"
	"popcorn/recommender/loader"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Postgres Driver
	"github.com/schollz/progressbar"
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

	loader.SetDatasetDir(viper.GetString("ml.dataset_dir"))
	if err := loader.LoadMovies(); err != nil {
		return err
	}

	if err := loader.LoadMetadata(); err != nil {
		return err
	}

	if err := loader.LoadTags(); err != nil {
		return err
	}

	if err := loader.LoadRatings(); err != nil {
		return err
	}

	loader.AddRatingStatsToMovies()

	movies := loader.Movies()
	bar := progressbar.New(len(movies))

	var count int
	for _, movie := range loader.Movies() {
		if err := db.Create(movie).Error; err != nil {
			logrus.Error(err)
		} else {
			bar.Add(1)
			count++
		}
	}

	fmt.Println()
	logrus.Infof("inserted %d movies to database", count)

	return nil
}
