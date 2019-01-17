package cmd

import (
	"fmt"
	"net"
	"os"
	"popcorn/recommender/pb/movie"
	"popcorn/recommender/rec"

	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// EnvConfig specifies environmental variable
type EnvConfig struct {
	GCP         bool `env:"GCP"          envDefault:"false"`
	DefaultPort int  `env:"DEFAULT_PORT" envDefault:"8081"`
	GCPPort     int  `env:"GCP_PORT"     envDefault:"8080"`
}

// Serve accepts incoming gRPC requests and handle them with registered services.
func Serve(cmd *cobra.Command, args []string) error {
	cfg := EnvConfig{}
	if err := env.Parse(&cfg); err != nil {
		return err
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
	movie.RegisterRecommendationServer(srv, &rec.Server{})

	logrus.Infof("recommender is listening and serving on port %d", port)
	return srv.Serve(lis)
}
