package cmd

import (
	"popcorn/recommender/loader"
	"popcorn/recommender/lowrank"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func initCSVLoaderForTraining() error {
	loader.SetDatasetDir(viper.GetString("data.dir"))
	if err := loader.LoadMovies(); err != nil {
		return err
	}

	if err := loader.LoadRatings(); err != nil {
		return err
	}

	return nil
}

// Train is a command for training the recommendation model.
func Train(cmd *cobra.Command, args []string) error {
	viper.SetConfigName("training")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := initCSVLoaderForTraining(); err != nil {
		return err
	}

	K := viper.GetInt("training.feature_dim")

	lowrank.SetMinRatingsPerUser(viper.GetInt("data.filter.min_ratings_per_user"))
	f, err := lowrank.NewIterativeFactorizer(K)
	if err != nil {
		return err
	}

	err = f.Train(
		viper.GetInt("training.num_steps"),
		viper.GetInt("training.epoch_size"),
		viper.GetFloat64("training.regularization"),
		viper.GetFloat64("training.learning_rate"),
	)

	if err != nil {
		return err
	}

	logrus.Info("completed training, now exporting to CSV")
	if err := loader.ExportMovieLatentVector(f.MovieFeatures(), K); err != nil {
		return err
	}

	logrus.Infof("feature CSV is saved")
	return nil
}
