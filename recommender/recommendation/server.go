package recommendation

import (
	"context"
	"fmt"
	"popcorn/recommender/lowrank"
	"popcorn/recommender/model"
	"popcorn/recommender/protogo"

	"github.com/sirupsen/logrus"
)

// Server provides gRPC service.
type Server struct{}

// Fetch returns a list of recommended movies.
func (srv *Server) Fetch(ctx context.Context, req *protogo.RecommendRequest) (*protogo.RecommendResponse, error) {
	logrus.Infof("server received request to fetch movie for user %s", req.UserEmail)

	ratings, err := model.FetchUserRatings(req.UserEmail)
	if err != nil {
		return nil, err
	}

	if len(ratings) == 0 {
		return nil, fmt.Errorf("cannot recommend movies for %s, he/she hasn't rated any movie", req.UserEmail)
	}

	var avg float64
	for _, rating := range ratings {
		avg += float64(rating.Value)
	}

	avg /= float64(len(ratings))

	pref, err := model.FetchUserPreference(req.UserEmail)
	if err != nil {
		return nil, err
	}

	movieIDs := []model.MovieID{}
	for movieID, feature := range movieFeatureStore {
		pred, err := lowrank.DotProduct(pref.Value, feature)
		if err != nil {
			logrus.Error(err)
			continue
		}

		if pred > 2.5 {
			movieIDs = append(movieIDs, movieID)
		}

		if len(movieIDs) == 10 {
			break
		}
	}

	fmt.Println(movieIDs)
	movies, err := model.FetchMovies(movieIDs)
	if err != nil {
		return nil, err
	}

	protoMovies := []*protogo.Movie{}
	for _, m := range movies {
		protoMovies = append(protoMovies, &protogo.Movie{
			Id:     int64(m.ID),
			Title:  m.Title,
			ImdbId: m.IMDBID,
		})
	}

	return &protogo.RecommendResponse{Movies: protoMovies}, nil
}

// UpdateUserPreference queues a training task.
func (srv *Server) UpdateUserPreference(ctx context.Context, req *protogo.UpdateRequest) (*protogo.UpdateResponse, error) {
	var latent []float64

	pref, err := model.FetchUserPreference(req.UserEmail)
	if err == nil {
		latent = pref.Value
	} else {
		latent = lowrank.RandVector(userLatentDim)
	}

	ratings, err := model.FetchUserRatings(req.UserEmail)
	if err != nil {
		return nil, err
	}

	ratingMap := make(map[model.MovieID]float64)
	for _, r := range ratings {
		ratingMap[r.MovieID] = float64(r.Value)
	}

	ch := make(chan lowrank.TrainerResponse)
	JobQueue <- lowrank.TrainerJob{
		UserEmail:      req.UserEmail,
		UserRatings:    ratingMap,
		UserPreference: latent,
		Response:       ch,
	}

	go handleResponse(req.UserEmail, ch)

	return &protogo.UpdateResponse{Accepted: true}, nil
}

func handleResponse(email string, ch <-chan lowrank.TrainerResponse) {
	resp := <-ch
	logrus.Infof("received trainer response for user %s", email)
	if resp.FinalLoss >= resp.InitLoss {
		logrus.Warn("training is not successful, final loss is greater than initial loss")
		return
	}

	if err := model.InsertUpdateUserPreference(email, resp.Preference); err != nil {
		logrus.Error(err)
	}
}
