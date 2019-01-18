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

	for i := 0; i < 10; i++ {
		detail := &model.MovieDetail{
			IMDBID: fmt.Sprintf("tttt011470%d", i),
			Detail: json.RawMessage(`{foo: "bar"}`),
		}

		if err := db.Create(detail).Error; err != nil {
			logrus.Error(err)
		}
	}

	return nil
}
