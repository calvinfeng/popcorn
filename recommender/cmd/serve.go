package cmd

import (
	"context"
	"fmt"
	"net"
	"os"
	"popcorn/recommender/model"
	pbmovie "popcorn/recommender/pb/movie"
	"popcorn/recommender/recommendation"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

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

	if err := model.ConnectDB(); err != nil {
		return err
	}

	if err := recommendation.InitStore(); err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := recommendation.RunTrainingGround(ctx); err != nil {
		return err
	}

	srv := grpc.NewServer()
	pbmovie.RegisterRecommendationServer(srv, &recommendation.Server{})

	logrus.Infof("recommender is listening and serving on port %d", viper.GetInt("grpc.port"))
	return srv.Serve(lis)
}
