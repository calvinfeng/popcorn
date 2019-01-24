package cmd

import (
	"context"
	"fmt"
	"net"
	"os"
	"popcorn/recommender/model"
	pbmovie "popcorn/recommender/pb/movie"
	"popcorn/recommender/recommendation"
	"runtime"

	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// EnvConfig captures environmental variable
type EnvConfig struct {
	GCP      bool `env:"GCP"    envDefault:"false"`
	Docker   bool `env:"DOCKER" envDefault:"false"`
	Training bool `env:"TRAIN"  envDefault:"false"`
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
	} else if cfg.Training {
		viper.SetConfigName("training")
	} else {
		viper.SetConfigName("development")
	}

	return viper.ReadInConfig()
}

// Serve accepts incoming gRPC requests and handle them with registered services.
func Serve(cmd *cobra.Command, args []string) error {
	// go func() {
	// 	for {
	// 		memUsage()
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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

	if err := recommendation.RunTrainingGround(ctx); err != nil {
		return err
	}

	srv := grpc.NewServer()
	pbmovie.RegisterRecommendationServer(srv, &recommendation.Server{})

	logrus.Infof("recommender is listening and serving on port %d", viper.GetInt("grpc.port"))
	return srv.Serve(lis)
}

// MemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
//
// For info on each, see: https://golang.org/pkg/runtime/#MemStats
func memUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
