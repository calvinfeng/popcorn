package recommendation

import (
	"context"
	"popcorn/recommender/lowrank"
	"popcorn/recommender/model"
)

var jobQueue = make(chan lowrank.TrainingJob)

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
		}
	}
}
