package main

import (
	"fmt"
	"net"
	"os"
	"popcorn/recommender/pb/movie"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const port = 8081

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	srv := grpc.NewServer()

	// Register services
	movie.RegisterRecommendationServer(srv, &RecommendationService{})

	// Launch server
	logrus.Infof("Golang gRPC server is listening and serving on port %d", port)
	if err := srv.Serve(lis); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
