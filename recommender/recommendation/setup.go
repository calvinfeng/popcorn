package recommendation

import (
	"popcorn/recommender/loader"
	"popcorn/recommender/lowrank"
	"popcorn/recommender/model"
)

var trainer *lowrank.Trainer

// InitTrainer pulls all movies from database and initialize a trainer with all the movie features
// information.
func InitTrainer() error {
	movies, err := model.AllMovies()
	if err != nil {
		return err
	}

	movieLatentMap := make(map[loader.MovieID][]float64)
	for _, movie := range movies {
		if len(movie.Feature) == 0 {
			continue
		}

		movieLatentMap[loader.MovieID(movie.ID)] = movie.Feature
	}

	trainer = lowrank.NewTrainer(movieLatentMap)

	return nil
}
