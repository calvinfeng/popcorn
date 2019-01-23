package recommendation

import (
	"context"
	"errors"
	"fmt"
	"popcorn/recommender/model"
	"popcorn/recommender/pb/movie"

	"github.com/sirupsen/logrus"
)

// Server provides gRPC service.
type Server struct{}

// Fetch returns a list of recommended movies.
func (srv *Server) Fetch(ctx context.Context, req *movie.RecommendRequest) (*movie.RecommendResponse, error) {
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
		pred, err := dotProduct(pref.Value, feature)
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

	protoMovies := []*movie.Movie{}
	for _, m := range movies {
		protoMovies = append(protoMovies, &movie.Movie{
			Id:     int64(m.ID),
			Title:  m.Title,
			ImdbId: m.IMDBID,
		})
	}

	return &movie.RecommendResponse{Movies: protoMovies}, nil
}

// UpdateUserPreference queues a training task.
func (srv *Server) UpdateUserPreference(ctx context.Context, req *movie.UpdateRequest) (*movie.UpdateResponse, error) {
	return &movie.UpdateResponse{Accepted: true}, nil
}

func dotProduct(v, u []float64) (product float64, err error) {
	if len(v) != len(u) {
		err = errors.New("vectors are different length")
		return
	}

	for i := 0; i < len(v); i++ {
		product += v[i] * u[i]
	}

	return
}
