package recommendation

import (
	"popcorn/recommender/lowrank"
	"popcorn/recommender/model"
)

var trainer *lowrank.Trainer

// InitTrainer pulls all movies from database and initialize a trainer with all the movie features
// information.
func InitTrainer() error {
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

	trainer = lowrank.NewTrainer(movieLatentMap)

	return nil
}
