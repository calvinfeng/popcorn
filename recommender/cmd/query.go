package cmd

import (
	"math/rand"
	"popcorn/recommender/model"

	_ "github.com/jinzhu/gorm/dialects/postgres" // Postgres Driver
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Query runs an example query on a PostgreSQL database.
func Query(cmd *cobra.Command, args []string) error {
	if err := configureViper(); err != nil {
		return err
	}

	if err := model.ConnectDB(); err != nil {
		return err
	}

	pref := make([]float64, 10)
	for i := 0; i < len(pref); i++ {
		pref[i] = rand.Float64()
	}

	if err := model.CreateUpdateUserPreference("cfeng@example.com", pref); err != nil {
		return err
	}

	if err := model.InsertUserRating("cfeng@example.com", 1, 4.5); err != nil {
		logrus.Error(err)
	}

	return nil
}
