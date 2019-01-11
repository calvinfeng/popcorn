package main

import (
	"net"
	"os"
	"popcorn/recommender/pb/movie"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	srv := grpc.NewServer()

	// Register services
	movie.RegisterRecommendationServer(srv, &RecommendationService{})

	// Launch server
	logrus.Info("gRPC server is serving port 8080")
	if err := srv.Serve(lis); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
