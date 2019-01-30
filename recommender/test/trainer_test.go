package test

import (
	"context"
	"io/ioutil"
	"math/rand"
	"popcorn/recommender/lowrank"
	"popcorn/recommender/model"
	"popcorn/recommender/recommendation"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("../conf")
	viper.SetConfigName("testing")
	logrus.SetOutput(ioutil.Discard)
}

func TestLowRankTrainer(t *testing.T) {
	err := viper.ReadInConfig()
	assert.NoError(t, err)

	err = model.ConnectDB()
	assert.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go recommendation.ProcessJob(ctx)

	movies, err := model.FetchAllMovies()
	assert.NoError(t, err)
	assert.True(t, len(movies) > 0)

	K := viper.GetInt("personalized_training.feature_dim")
	lowrank.UpdateTrainerConfig(
		viper.GetInt("personalized_training.num_steps"),
		viper.GetFloat64("personalized_training.regularization"),
		viper.GetFloat64("personalized_training.learning_rate"),
	)

	t.Run("SubmitTrainingJob", func(t *testing.T) {
		rand.Seed(1)

		ratings := make(map[model.MovieID]float64)
		for len(ratings) < 10 {
			movie := movies[rand.Intn(len(movies))]
			ratings[movie.ID] = 5
		}

		pref := make([]float64, K)
		for k := 0; k < K; k++ {
			pref[k] = rand.Float64()
		}

		resp := make(chan lowrank.TrainerResponse)
		job := lowrank.TrainerJob{
			UserEmail:      "calvin@example.com",
			UserRatings:    ratings,
			UserPreference: pref,
			Response:       resp,
		}

		recommendation.JobQueue <- job
		trainerResp := <-resp

		t.Logf("init_loss=%f and final_loss=%f", trainerResp.InitLoss, trainerResp.FinalLoss)
		t.Logf("%v", trainerResp.Preference)

		assert.Len(t, trainerResp.Preference, K)
		assert.True(t, trainerResp.FinalLoss < trainerResp.InitLoss)
	})
}
