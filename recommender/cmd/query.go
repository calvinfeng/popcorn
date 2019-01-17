package cmd

import (
	"encoding/json"
	"fmt"
	"popcorn/recommender/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Postgres Driver
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Query runs an example query on a PostgreSQL database.
func Query(cmd *cobra.Command, args []string) error {
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

	movies := []*model.Movie{}
	if err := db.Find(&movies).Error; err != nil {
		return err
	}

	fmt.Println(movies)

	detail := &model.MovieDetail{
		IMDBID: "tttt0114709",
		Detail: json.RawMessage(`{foo: "bar"}`),
	}

	if err := db.Create(detail).Error; err != nil {
		logrus.Error(err)
	}

	detail = &model.MovieDetail{
		IMDBID: "tttt0114709",
	}

	if err := db.Find(&detail).Error; err != nil {
		return err
	}

	fmt.Println("Found detail for", detail.IMDBID, string(detail.Detail))

	return nil
}
