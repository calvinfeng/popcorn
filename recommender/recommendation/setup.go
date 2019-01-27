package recommendation

import (
	"context"
	"math"
	"popcorn/recommender/lowrank"
	"popcorn/recommender/model"

	"github.com/sirupsen/logrus"
)

// JobQueue is a queue for personalized training.
var JobQueue = make(chan lowrank.TrainerJob)

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
		case job := <-JobQueue:
			trainer := lowrank.NewTrainer(movieLatentMap)
			trainer.AssignJob(job)

			go func(t *lowrank.Trainer) {
				initLoss, err := t.Loss()
				if err != nil {
					logrus.Error(err)
					return
				}

				err = t.Train()
				if err != nil {
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

				finalLoss, err := t.Loss()
				if err != nil {
					logrus.Error(err)
					return
				}

				job.Response <- lowrank.TrainerResponse{
					InitLoss:   initLoss,
					FinalLoss:  finalLoss,
					Preference: t.Preference(),
				}
			}(trainer)
		}
	}
}
