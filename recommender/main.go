package main

import (
	"fmt"
	"net"
	"os"
	"popcorn/recommender/pb/movie"

	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// EnvConfig specifies environmental variable
type EnvConfig struct {
	GCP         bool `env:"GCP"          envDefault:"false"`
	DefaultPort int  `env:"DEFAULT_PORT" envDefault:"8081"`
	GCPPort     int  `env:"GCP_PORT"     envDefault:"8080"`
}

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
}

func main() {
	cfg := EnvConfig{}
	err := env.Parse(&cfg)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	port := cfg.DefaultPort
	if cfg.GCP {
		port = cfg.GCPPort
	}

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
