package test

import (
	"io/ioutil"
	"math"
	"math/rand"
	"popcorn/recommender/loader"
	"popcorn/recommender/lowrank"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func init() {
	viper.AddConfigPath("../conf")
	viper.SetConfigName("testing")
	logrus.SetOutput(ioutil.Discard)
}

func initCSVLoaderForTesting() error {
	loader.SetDatasetDir(viper.GetString("data.dir"))
	if err := loader.LoadMovies(); err != nil {
		return err
	}

	if err := loader.LoadRatings(); err != nil {
		return err
	}

	return nil
}

func TestTrainingGradientCalculation(t *testing.T) {
	err := viper.ReadInConfig()
	assert.NoError(t, err)

	err = initCSVLoaderForTesting()
	assert.NoError(t, err)

	rand.Seed(0)

	f, err := lowrank.NewIterativeFactorizer(viper.GetInt("model_training.feature_dim"))
	assert.NoError(t, err)

	users := f.Users()
	for i := 0; i < 10; i++ {
		userID := users[rand.Intn(len(users))]
		disp, err := lowrank.UserLatentGradientCheck(f, userID, 0.5, 0.001)
		assert.NoError(t, err)
		assert.NotEmpty(t, disp)
		assertAlmostZero(t, disp, 5)
	}

	movies := f.Movies()
	for i := 0; i < 10; i++ {
		movieID := movies[rand.Intn(len(movies))]
		disp, err := lowrank.MovieLatentGradientCheck(f, movieID, 0.5, 0.001)
		assert.NoError(t, err)
		assert.NotEmpty(t, disp)
		assertAlmostZero(t, disp, 5)
	}
}

func assertAlmostZero(t *testing.T, actual []float64, precision int) {
	expected := make([]float64, len(actual))
	for i := 0; i < len(expected); i++ {
		assert.Equal(t,
			math.Round(expected[i]*math.Pow(10, float64(precision)))/math.Pow(10, float64(precision)),
			math.Round(actual[i]*math.Pow(10, float64(precision)))/math.Pow(10, float64(precision)),
		)
	}
}

func assertAlmostEqual(t *testing.T, expected, actual []float64, precision int) {
	assert.Equal(t, len(expected), len(actual))
	for i := 0; i < len(expected); i++ {
		assert.Equal(t,
			math.Round(expected[i]*math.Pow(10, float64(precision)))/math.Pow(10, float64(precision)),
			math.Round(actual[i]*math.Pow(10, float64(precision)))/math.Pow(10, float64(precision)),
		)
	}
}
