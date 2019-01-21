package cmd

import (
	"popcorn/recommender/loader"
	"popcorn/recommender/lowrank"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Train is a command for training the recommendation model.
func Train(cmd *cobra.Command, args []string) error {
	if err := configureViper(); err != nil {
		return err
	}

	lowrank.SetMinRatingsPerUser(viper.GetInt("min_ratings_per_user"))

	f, err := lowrank.NewIterativeFactorizer(viper.GetString("ml.dataset_dir"),
		viper.GetInt("ml.feature_dim"))
	if err != nil {
		return err
	}

	err = f.Train(
		viper.GetInt("ml.steps"),
		viper.GetInt("ml.epoch"),
		viper.GetFloat64("ml.regularization"),
		viper.GetFloat64("ml.learning_rate"),
	)

	if err != nil {
		return err
	}

	logrus.Info("completed training, now exporting to CSV")
	return loader.ExportMovieLatentVector(f.MovieFeatures(), viper.GetInt("ml.feature_dim"))
}
