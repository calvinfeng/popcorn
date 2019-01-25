package cmd

import (
	"fmt"
	"runtime"

	"github.com/caarlos0/env"
	"github.com/spf13/viper"
)

type envConfig struct {
	GCP bool `env:"GCP"    envDefault:"false"`
}

func configureViper() error {
	cfg := envConfig{}
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
