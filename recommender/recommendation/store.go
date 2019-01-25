package recommendation

import (
	"errors"
	"popcorn/recommender/model"

	"github.com/sirupsen/logrus"
)

var movieFeatureStore = make(map[model.MovieID][]float64)

// InitStore populates movie store with all movie features from database.
func InitStore() error {
	if !model.IsDatabaseConnected() {
		return errors.New("database is not connected")
	}

	movies, err := model.FetchAllMovies()
	if err != nil {
		return err
	}

	for _, movie := range movies {
		if len(movie.Feature) > 0 {
			movieFeatureStore[movie.ID] = movie.Feature
		}
	}

	logrus.Infof("movie feature store has been initialized with %d movies", len(movieFeatureStore))

	return nil
}
