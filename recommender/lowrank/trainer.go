package lowrank

import (
	"errors"
	"math"
	"popcorn/recommender/model"
)

// UpdateTrainerConfig is a setter for the local configuration variable.
func UpdateTrainerConfig(steps int, reg, learnRate float64) {
	cfg.Steps = steps
	cfg.Regularization = reg
	cfg.LearningRate = learnRate
}

// Default configuration value
var cfg = TrainerConfig{
	Steps:          1000,
	Regularization: 0.05,
	LearningRate:   2e-5,
}

// TrainerConfig configures the training hyperparamter for a trainer.
type TrainerConfig struct {
	Steps          int
	Regularization float64
	LearningRate   float64
}

// TrainerJob carries the payload necessary to perform an update on user preference.
type TrainerJob struct {
	UserEmail      string
	UserRatings    map[model.MovieID]float64
	UserPreference []float64
	Response       chan TrainerResponse
}

// TrainerResponse reports the training result from a trainer.
type TrainerResponse struct {
	InitLoss   float64
	FinalLoss  float64
	Preference []float64
}

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

// AssignJob assigns a training job to trainer.
func (t *Trainer) AssignJob(j TrainerJob) {
	t.userRatings = j.UserRatings
	t.userPreference = j.UserPreference
}

// Preference returns an updated preference of a trained user.
func (t *Trainer) Preference() []float64 {
	return t.userPreference
}

// Train performs training on a user.
func (t *Trainer) Train() error {
	for i := 0; i < cfg.Steps; i++ {
		grad, err := t.userLatentGradient(cfg.Regularization)
		if err != nil {
			return err
		}

		for k := 0; k < len(t.userPreference); k++ {
			t.userPreference[k] -= cfg.LearningRate * grad[k]
		}
	}
	return nil
}

// Loss returns the current loss.
func (t *Trainer) Loss() (loss float64, err error) {
	var count int
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
		count++
	}

	loss /= float64(count)

	for _, val := range t.userPreference {
		loss += 0.5 * cfg.Regularization * math.Pow(val, 2)
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
