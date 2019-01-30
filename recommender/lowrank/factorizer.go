package lowrank

import (
	"math"
	"math/rand"
	"popcorn/recommender/loader"
	"popcorn/recommender/model"
	"time"

	"github.com/sirupsen/logrus"
)

// NewIterativeFactorizer returns a new instance of iterative factorizer.
func NewIterativeFactorizer(K int) (*IterativeFactorizer, error) {
	f := &IterativeFactorizer{
		userLatentMap:  make(map[model.UserID][]float64),
		movieLatentMap: make(map[model.MovieID][]float64),
		userRating:     make(map[model.UserID]map[model.MovieID]float64),
		movieRating:    make(map[model.MovieID]map[model.UserID]float64),
		testSet:        make(map[model.UserID]map[model.MovieID]float64),
	}

	var trainCount, testCount int

	ratings := loader.RatingsFilteredByCount(minRatingsPerUser)
	for userID := range ratings {
		f.userLatentMap[userID] = RandVector(K)
		f.userRating[userID] = make(map[model.MovieID]float64)
		f.testSet[userID] = make(map[model.MovieID]float64)

		for movieID := range ratings[userID] {
			if _, ok := f.movieLatentMap[movieID]; !ok {
				f.movieLatentMap[movieID] = RandVector(K)
				f.movieRating[movieID] = make(map[model.UserID]float64)
			}

			if rand.Float64() < trainTestRatio {
				f.testSet[userID][movieID] = ratings[userID][movieID]
				testCount++
			} else {
				f.userRating[userID][movieID] = ratings[userID][movieID]
				f.movieRating[movieID][userID] = ratings[userID][movieID]
				trainCount++
			}
		}
	}

	logrus.Infof("factorizer has been initialized with %d training examples and %d test examples",
		trainCount, testCount)

	return f, nil
}

// IterativeFactorizer is optimized for sparsed matrix factorization. Instead of having a matrix
// with mostly empty values, it uses a map of user ID to map of movie ID to rating, and a map of
// movie ID to a map of user ID to movie ID.
type IterativeFactorizer struct {
	// Result of factorization
	userLatentMap  map[model.UserID][]float64
	movieLatentMap map[model.MovieID][]float64

	// Map of user ID to a map of movieID to rating submitted by the user.
	userRating map[model.UserID]map[model.MovieID]float64

	// Map of movie ID to a map of user ID to rating submitted by the user.
	movieRating map[model.MovieID]map[model.UserID]float64

	// Validation/Test set which is user-centric.
	testSet map[model.UserID]map[model.MovieID]float64
}

// MovieFeatures returns a map of movie ID to its feature vector.
func (f *IterativeFactorizer) MovieFeatures() map[model.MovieID][]float64 {
	return f.movieLatentMap
}

// Users return the list of all user ID(s) in factorizer.
func (f *IterativeFactorizer) Users() []model.UserID {
	ids := make([]model.UserID, 0, len(f.userLatentMap))
	for id := range f.userLatentMap {
		ids = append(ids, id)
	}

	return ids
}

// Movies return the list of all movie ID(s) in factorizer.
func (f *IterativeFactorizer) Movies() []model.MovieID {
	ids := make([]model.MovieID, 0, len(f.movieLatentMap))
	for id := range f.movieLatentMap {
		ids = append(ids, id)
	}

	return ids
}

// Loss returns the loss and validation root mean square error from current training result.
func (f *IterativeFactorizer) Loss(reg float64) (loss, rmse float64, err error) {
	for userID := range f.userRating {
		u := f.userLatentMap[userID]
		for movieID := range f.userRating[userID] {
			v := f.movieLatentMap[movieID]

			pred, err := DotProduct(u, v)
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
			pred, err := DotProduct(f.userLatentMap[userID], f.movieLatentMap[movieID])
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

// Train performs gradient descent training on latent vectors of users and movies.
func (f *IterativeFactorizer) Train(N, epochSize int, reg, learnRate float64) error {
	for i := 0; i < N; i++ {
		if i%epochSize == 0 {
			loss, rmse, err := f.Loss(reg)
			if err != nil {
				return err
			}

			logrus.Infof("%3d/%3d loss = %5.2f & rmse = %1.8f", i, N, loss, rmse)
		}

		start := time.Now()

		var err error
		uGrad := make(map[model.UserID][]float64)
		for userID := range f.userLatentMap {
			uGrad[userID], err = f.userLatentGradient(userID, reg)
			if err != nil {
				return err
			}
		}

		mGrad := make(map[model.MovieID][]float64)
		for movieID := range f.movieLatentMap {
			mGrad[movieID], err = f.movieLatentGradient(movieID, reg)
			if err != nil {
				return err
			}
		}

		// Perform latent vector updates, only after all gradients have been computed.
		for userID := range f.userLatentMap {
			for k := 0; k < len(f.userLatentMap[userID]); k++ {
				f.userLatentMap[userID][k] -= learnRate * uGrad[userID][k]
			}
		}

		for movieID := range f.movieLatentMap {
			for k := 0; k < len(f.movieLatentMap[movieID]); k++ {
				f.movieLatentMap[movieID][k] -= learnRate * mGrad[movieID][k]
			}
		}

		logrus.Printf("per training elapsed time: %s\n", time.Since(start))
	}

	return nil
}

func (f *IterativeFactorizer) userLatentGradient(id model.UserID, reg float64) ([]float64, error) {
	latent := f.userLatentMap[id]

	grad := make([]float64, len(latent))
	for k := 0; k < len(latent); k++ {
		for movieID := range f.userRating[id] {
			pred, err := DotProduct(latent, f.movieLatentMap[movieID])
			if err != nil {
				return nil, err
			}

			grad[k] += -1 * (f.userRating[id][movieID] - pred) * f.movieLatentMap[movieID][k]
		}

		grad[k] += reg * latent[k]
	}

	return grad, nil
}

func (f *IterativeFactorizer) movieLatentGradient(id model.MovieID, reg float64) ([]float64, error) {
	latent := f.movieLatentMap[id]

	grad := make([]float64, len(latent))
	for k := 0; k < len(latent); k++ {
		for userID := range f.movieRating[id] {
			pred, err := DotProduct(latent, f.userLatentMap[userID])
			if err != nil {
				return nil, err
			}

			grad[k] += -1 * (f.movieRating[id][userID] - pred) * f.userLatentMap[userID][k]
		}

		grad[k] += reg * latent[k]
	}

	return grad, nil
}
