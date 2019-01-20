package lowrank

import (
	"math"
	"popcorn/recommender/loader"
	"popcorn/recommender/model"
)

// IterativeFactorizer is optimized for sparsed matrix factorization. Instead of having a matrix
// with mostly empty values, it uses a map of user ID to map of movie ID to rating, and a map of
// movie ID to a map of user ID to movie ID.
type IterativeFactorizer struct {
	movies map[loader.MovieID]*model.Movie

	// Result of factorization
	userLatentMap  map[loader.UserID][]float64
	movieLatentMap map[loader.MovieID][]float64

	// userRating is a map of user ID to a map of movieID to rating submitted by the user.
	userRating map[loader.UserID]map[loader.MovieID]float64

	// movieRating is a map of movie ID to a map of user ID to rating submitted by the user.
	movieRating map[loader.MovieID]map[loader.UserID]float64

	// Validation/Test phase
	testSet map[loader.UserID]map[loader.MovieID]float64
}

// Loss returns the loss and validation root mean square error from current training result.
func (f *IterativeFactorizer) Loss(reg float64) (loss, rmse float64, err error) {
	for userID := range f.userRating {
		u := f.userLatentMap[userID]
		for movieID := range f.userRating[userID] {
			v := f.movieLatentMap[movieID]

			pred, err := dotProduct(u, v)
			if err != nil {
				return loss, rmse, err
			}

			loss += 0.5 * math.Pow(f.userRating[userID][movieID]-pred, 2)
		}
	}

	// User latent regularization
	for userID := range f.userLatentMap {
		for _, val := range f.userLatentMap[userID] {
			loss += 0.5 * reg * math.Pow(val, 2)
		}
	}

	// Movie latent regularization
	for movieID := range f.movieLatentMap {
		for _, val := range f.movieLatentMap[movieID] {
			loss += 0.5 * reg * math.Pow(val, 2)
		}
	}

	var count int
	for userID := range f.testSet {
		for movieID := range f.testSet[userID] {
			pred, err := dotProduct(f.userLatentMap[userID], f.movieLatentMap[movieID])
			if err != nil {
				return loss, rmse, err
			}

			count++
			rmse += math.Pow(f.testSet[userID][movieID]-pred, 2)
		}
	}

	rmse = math.Sqrt(rmse / float64(count))

	return
}

func (f *IterativeFactorizer) userLatentGradient(id loader.UserID, reg float64) ([]float64, error) {
	latent := f.userLatentMap[id]

	grad := make([]float64, len(latent))
	for k := 0; k < len(latent); k++ {
		for movieID := range f.userRating[id] {
			pred, err := dotProduct(latent, f.movieLatentMap[movieID])
			if err != nil {
				return nil, err
			}

			grad[k] += -1 * (f.userRating[id][movieID] - pred) * f.movieLatentMap[movieID][k]
		}

		grad[k] += reg * latent[k]
	}

	return grad, nil
}

func (f *IterativeFactorizer) movieLatentGradient(id loader.MovieID, reg float64) ([]float64, error) {
	latent := f.movieLatentMap[id]

	grad := make([]float64, len(latent))
	for k := 0; k < len(latent); k++ {
		for userID := range f.movieRating[id] {
			pred, err := dotProduct(latent, f.userLatentMap[userID])
			if err != nil {
				return nil, err
			}

			grad[k] += -1 * (f.movieRating[id][userID] - pred) * f.userLatentMap[userID][k]
		}

		grad[k] += reg * latent[k]
	}

	return grad, nil
}
