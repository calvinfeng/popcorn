package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var db *gorm.DB

// ConnectDB dials Postgres using settings from conf file.
func ConnectDB() (err error) {
	addr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.hostname"),
		viper.GetInt("database.port"),
		viper.GetString("database.name"),
		viper.GetString("database.ssl_mode"),
	)

	db, err = gorm.Open("postgres", addr)

	return
}

// IsDatabaseConnected checks if model package has already connected to database.
func IsDatabaseConnected() bool {
	return db != nil
}
