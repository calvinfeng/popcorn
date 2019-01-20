package cmd

import (
	"fmt"
	"math"
	"popcorn/recommender/loader"
	"popcorn/recommender/model"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Group performs nearest neighbor search to verify recommendation result.
func Group(cmd *cobra.Command, args []string) error {
	if err := configureViper(); err != nil {
		return err
	}

	loader.SetDatasetDir(viper.GetString("ml.dataset_dir"))
	if err := loader.LoadMovies(); err != nil {
		return err
	}

	if err := loader.LoadFeatures(); err != nil {
		return err
	}

	selection := map[uint]struct{}{
		318:   struct{}{},
		3578:  struct{}{},
		6377:  struct{}{},
		79132: struct{}{},
		60069: struct{}{},
	}

	movies := loader.Movies()
	for _, movie := range movies {
		if _, ok := selection[movie.ID]; ok && len(movie.Feature) > 0 {
			fmt.Printf("%s - %d has the following neighbors\n", movie.Title, movie.Year)
			nn := nearest(movies, movie, 10)
			for _, neighbor := range nn {
				fmt.Printf("\t%s - %d\n", neighbor.Title, neighbor.Year)
			}
		}
	}

	return nil
}

// Todo: use min heap if performance is a problem.
func nearest(movies []*model.Movie, m *model.Movie, N int) map[uint]*model.Movie {
	nn := make(map[uint]*model.Movie)

	for i := 0; i < N; i++ {
		nearest := movies[0]
		if nearest.ID == m.ID {
			nearest = movies[1] // Just in case we selected the first movie
		}

		minDist := distance(m.Feature, nearest.Feature)

		for _, other := range movies {
			if len(other.Feature) == 0 {
				continue
			}

			if _, ok := nn[other.ID]; ok || other.ID == m.ID {
				continue
			}

			dist := distance(m.Feature, other.Feature)
			if dist < minDist {
				nearest = other
				minDist = dist
			}
		}

		nn[nearest.ID] = nearest
	}

	return nn
}

func distance(u, v []float64) float64 {
	if len(u) != len(v) {
		panic("vectors have different length")
	}

	var dist float64
	for i := 0; i < len(u); i++ {
		dist += math.Pow(u[i]-v[i], 2)
	}

	return math.Sqrt(dist)
}
