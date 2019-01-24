package recommendation

import (
	"context"
	"math"
	"popcorn/recommender/lowrank"
	"popcorn/recommender/model"

	"github.com/sirupsen/logrus"
)

var jobQueue = make(chan lowrank.TrainingJob)

// TODO: Make this a config
var trainingSteps = 1000
var regularization = 0.05
var learnRate = 0.00001

// RunTrainingGround kicks off a goroutine to listen for new training job.
func RunTrainingGround(ctx context.Context) error {
	movies, err := model.FetchAllMovies()
	if err != nil {
		return err
	}

	movieLatentMap := make(map[model.MovieID][]float64)
	for _, movie := range movies {
		if len(movie.Feature) == 0 {
			continue
		}

		movieLatentMap[movie.ID] = movie.Feature
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case job := <-jobQueue:
			trainer := lowrank.NewTrainer(movieLatentMap)
			trainer.AssignJob(job)

			go func(t *lowrank.Trainer) {
				if err := t.Train(trainingSteps, regularization, learnRate); err != nil {
					logrus.Error(err)
					return
				}

				logrus.Infof("completed training for user %s", job.UserEmail)
				for _, el := range t.Preference() {
					if math.IsNaN(el) || math.IsInf(el, 0) {
						logrus.Errorf("user %s updated preference has NaN / Inf value", job.UserEmail)
						return
					}
				}

				if err := model.InsertUpdateUserPreference(job.UserEmail, t.Preference()); err != nil {
					logrus.Error(err)
				}
			}(trainer)
		}
	}
}
