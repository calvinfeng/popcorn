package lowrank

import (
	"math"
	"popcorn/recommender/loader"
)

// UserLatentGradientCheck performs gradient check on a factorizer's gradient computation for user
// latent vector.
func UserLatentGradientCheck(f *IterativeFactorizer, id loader.UserID, reg, h float64) ([]float64, error) {
	grad, err := f.userLatentGradient(id, reg)
	if err != nil {
		return nil, err
	}

	discrepancy := make([]float64, len(grad))
	for k := 0; k < len(grad); k++ {
		val := f.userLatentMap[id][k]

		// f(x+h)
		f.userLatentMap[id][k] = val + h
		fxph, _, err := f.Loss(reg)
		if err != nil {
			return nil, err
		}

		// f(x-h)
		f.userLatentMap[id][k] = val - h
		fxmh, _, err := f.Loss(reg)
		if err != nil {
			return nil, err
		}

		f.userLatentMap[id][k] = val
		numGrad := (fxph - fxmh) / (2 * h)

		discrepancy[k] = math.Abs(numGrad - grad[k])
	}

	return discrepancy, nil
}

// MovieLatentGradientCheck performs gradient check on a factorizer's gradient computation for movie
// latent vector.
func MovieLatentGradientCheck(f *IterativeFactorizer, id loader.MovieID, reg, h float64) ([]float64, error) {
	grad, err := f.movieLatentGradient(id, reg)
	if err != nil {
		return nil, err
	}

	discrepancy := make([]float64, len(grad))
	for k := 0; k < len(grad); k++ {
		val := f.movieLatentMap[id][k]

		// f(x+h)
		f.movieLatentMap[id][k] = val + h
		fxph, _, err := f.Loss(reg)
		if err != nil {
			return nil, err
		}

		// f(x-h)
		f.movieLatentMap[id][k] = val - h
		fxmh, _, err := f.Loss(reg)
		if err != nil {
			return nil, err
		}

		f.movieLatentMap[id][k] = val
		numGrad := (fxph - fxmh) / (2 * h)

		discrepancy[k] = math.Abs(numGrad - grad[k])
	}

	return discrepancy, nil
}
