package recommendation

import (
	"context"
	"popcorn/recommender/pb/movie"

	"github.com/sirupsen/logrus"
)

// Server provides gRPC service.
type Server struct{}

// Fetch returns a list of recommended movies.
func (srv *Server) Fetch(ctx context.Context, req *movie.RecommendRequest) (*movie.RecommendResponse, error) {
	logrus.Infof("server received request to fetch movie for user %s", req.UserEmail)

	movies := []*movie.Movie{
		&movie.Movie{
			Id:              1,
			Title:           "Toy Story",
			ImdbId:          "imdb1234",
			PredictedRating: 5.0,
		},
	}

	return &movie.RecommendResponse{Movies: movies}, nil
}

// UpdateUserPreference queues a training task.
func (srv *Server) UpdateUserPreference(ctx context.Context, req *movie.UpdateRequest) (*movie.UpdateResponse, error) {
	return &movie.UpdateResponse{Accepted: true}, nil
}
