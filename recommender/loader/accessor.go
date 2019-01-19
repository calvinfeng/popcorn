package loader

import "popcorn/recommender/model"

var dataDir string

// SetDatasetDir configures a dataset directory for which to find CSV data from.
func SetDatasetDir(dir string) {
	dataDir = dir
}

var movies = make(map[movieID]*model.Movie)

// Movies return the list of loaded movies.
func Movies() []*model.Movie {
	list := make([]*model.Movie, 0, len(movies))
	for _, movie := range movies {
		list = append(list, movie)
	}

	return list
}

type (
	movieID uint
	userID  uint
)

// Ratings is a map of user ID to a map of movie ID to rating value, range from 0 to 5.
var ratings = make(map[userID]map[movieID]float64)

// RatingsFilteredByCount returns ratings by user who has rated at least n movies.
func RatingsFilteredByCount(n int) map[userID]map[movieID]float64 {
	result := make(map[userID]map[movieID]float64)
	for k, v := range ratings {
		if len(v) >= n {
			result[k] = v
		}
	}

	return result
}
