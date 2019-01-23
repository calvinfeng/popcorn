package cmd

import (
	"fmt"
	"popcorn/recommender/model"
	"runtime"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres" // Postgres Driver
	"github.com/spf13/cobra"
)

// Query runs an example query on a PostgreSQL database.
func Query(cmd *cobra.Command, args []string) error {
	if err := configureViper(); err != nil {
		return err
	}

	if err := model.ConnectDB(); err != nil {
		return err
	}

	movies, err := model.FetchAllMovies()
	if err != nil {
		return err
	}

	memUsage()

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Millisecond)
	}

	var count int
	start := time.Now()
	for i := 0; i < len(movies); i++ {
		count++
	}

	fmt.Printf("pulled %d movies from database\n", count)
	fmt.Printf("elapsed %s\n", time.Since(start))

	return nil
}

// MemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
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
