package lowrank

import (
	"errors"
	"math"
	"popcorn/recommender/model"
)

// NewTrainer returns a trainer with movie latent map intialized.
func NewTrainer(latents map[model.MovieID][]float64) *Trainer {
	return &Trainer{
		movieLatentMap: latents,
	}
}

// Trainer is a partial matrix factorizer. It is meant to only update user latent preference during
// training. This trainer is used in runtime. During initial startup, this trainer will be initialized
// with pretrained movie features. Whenever frontend server makes a gRPC request to update user
// preference, this trainer will train the user.
type Trainer struct {
	userRatings    map[model.MovieID]float64
	userPreference []float64

	movieLatentMap map[model.MovieID][]float64
}

// SetUser assigns a user to trainer.
func (t *Trainer) SetUser(id model.UserID, ratings map[model.MovieID]float64, pref []float64) {
	t.userRatings = ratings
	t.userPreference = pref
}

// Preference returns an updated preference of a trained user.
func (t *Trainer) Preference() []float64 {
	return t.userPreference
}

// Train performs training on a user.
func (t *Trainer) Train(steps int, reg, learnRate float64) error {
	for i := 0; i < steps; i++ {
		grad, err := t.userLatentGradient(reg)
		if err != nil {
			return err
		}

		for k := 0; k < len(t.userPreference); k++ {
			t.userPreference[k] -= learnRate * grad[k]
		}
	}
	return nil
}

// Loss returns the current loss.
func (t *Trainer) Loss(reg float64) (loss float64, err error) {
	for movieID := range t.userRatings {
		if _, ok := t.movieLatentMap[movieID]; !ok {
			continue
		}

		var pred float64
		pred, err = dotProduct(t.userPreference, t.movieLatentMap[movieID])
		if err != nil {
			return
		}

		loss += 0.5 * math.Pow(t.userRatings[movieID]-pred, 2)
	}

	return
}

func (t *Trainer) userLatentGradient(reg float64) ([]float64, error) {
	if t.userRatings == nil || t.userPreference == nil {
		return nil, errors.New("there is no user for trainer to train")
	}

	grad := make([]float64, len(t.userPreference))
	for k := 0; k < len(t.userPreference); k++ {
		for movieID := range t.userRatings {
			if _, ok := t.movieLatentMap[movieID]; !ok {
				continue
			}

			pred, err := dotProduct(t.userPreference, t.movieLatentMap[movieID])
			if err != nil {
				return nil, err
			}

			grad[k] += -1 * (t.userRatings[movieID] - pred) * t.movieLatentMap[movieID][k]
		}

		grad[k] += reg * t.userPreference[k]
	}

	return grad, nil
}
