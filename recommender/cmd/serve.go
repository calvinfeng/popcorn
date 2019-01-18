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
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// EnvConfig captures environmental variable
type EnvConfig struct {
	GCP    bool `env:"GCP"    envDefault:"false"`
	Docker bool `env:"DOCKER" envDefault:"false"`
}

func init() {
	viper.AddConfigPath("./conf")
}

func configureViper() error {
	cfg := EnvConfig{}
	if err := env.Parse(&cfg); err != nil {
		return err
	}

	if cfg.GCP {
		viper.SetConfigName("production")
	} else {
		viper.SetConfigName("development")
	}

	return viper.ReadInConfig()
}

// Serve accepts incoming gRPC requests and handle them with registered services.
func Serve(cmd *cobra.Command, args []string) error {
	if err := configureViper(); err != nil {
		return err
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt("grpc.port")))
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	srv := grpc.NewServer()
	movie.RegisterRecommendationServer(srv, &rec.Server{})

	logrus.Infof("recommender is listening and serving on port %d", viper.GetInt("grpc.port"))
	return srv.Serve(lis)
}
