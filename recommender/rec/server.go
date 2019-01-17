package rec

import (
	"context"
	"popcorn/recommender/pb/movie"

	"github.com/sirupsen/logrus"
)

// Server provides gRPC service.
type Server struct{}

// Fetch returns a list of recommended movies.
func (srv *Server) Fetch(ctx context.Context, req *movie.RecRequest) (*movie.RecResponse, error) {
	logrus.Infof("server received request to fetch movie for user %d", req.UserId)

	movies := []*movie.Movie{
		&movie.Movie{
			Id:              1,
			Title:           "Toy Story",
			ImdbId:          "imdb1234",
			PredictedRating: 5.0,
		},
	}

	return &movie.RecResponse{Movies: movies}, nil
}
