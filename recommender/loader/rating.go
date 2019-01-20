package loader

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Oldest timestamp of rating the program is willing to accept.
var cutoffTimestamp int64 = 1167609600

// LoadRatings loads movie rating data from CSV file into memory.
func LoadRatings() error {
	csvFile, err := os.Open(fmt.Sprintf("%s/ratings.csv", dataDir))
	if err != nil {
		return err
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		row, err := reader.Read()
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		var uid, mid, timestamp int64
		var rating float64

		uid, err = strconv.ParseInt(row[0], 10, 64)
		if err != nil {
			continue
		}

		mid, err = strconv.ParseInt(row[1], 10, 64)
		if err != nil {
			continue
		}

		rating, err = strconv.ParseFloat(row[2], 64)
		if err != nil {
			continue
		}

		timestamp, err = strconv.ParseInt(row[3], 10, 64)
		if err != nil {
			continue
		}

		// Check if rating is submitted way too long ago
		if timestamp < cutoffTimestamp {
			continue
		}

		if _, ok := ratings[UserID(uid)]; !ok {
			ratings[UserID(uid)] = make(map[MovieID]float64)
		}

		ratings[UserID(uid)][MovieID(mid)] = rating
	}

	return nil
}

// AddRatingStatsToMovies calculates number of ratings each movie received and average rating from
// every user that rated the movie.
func AddRatingStatsToMovies() {
	for _, rated := range ratings {
		for MovieID, val := range rated {
			if _, ok := movies[MovieID]; !ok {
				continue
			}

			movies[MovieID].AverageRating += float32(val)
			movies[MovieID].NumRating++
		}
	}

	for _, movie := range movies {
		if (movie.NumRating) > 0 {
			movie.AverageRating /= float32(movie.NumRating)
		}
	}
}
