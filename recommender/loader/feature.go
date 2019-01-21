package loader

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"popcorn/recommender/model"
	"strconv"
)

// ExportMovieLatentVector saves movie features to a CSV file.
func ExportMovieLatentVector(features map[model.MovieID][]float64, K int) error {
	csvFile, err := os.Create(fmt.Sprintf("%s/features.csv", dataDir))
	if err != nil {
		return err
	}

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	header := []string{"movieId"}
	for k := 1; k <= K; k++ {
		header = append(header, fmt.Sprintf("f%d", k))
	}

	err = writer.Write(header)
	if err != nil {
		return err
	}

	for movieID, vector := range features {
		row := []string{strconv.Itoa(int(movieID))}
		for _, el := range vector {
			row = append(row, strconv.FormatFloat(el, 'f', 6, 64))
		}

		err = writer.Write(row)
		if err != nil {
			return err
		}
	}

	return nil
}

// LoadFeatures loads latent vectors from CSV files and insert them to movies as features.
func LoadFeatures() error {
	csvFile, err := os.Open(fmt.Sprintf("%s/features.csv", dataDir))
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

		var id int64
		id, err = strconv.ParseInt(row[0], 10, 64)
		if err != nil {
			continue
		}

		movie, ok := movies[model.MovieID(id)]
		if !ok {
			continue
		}

		movie.Feature = make([]float64, 0, len(row)-1)
		for k := 1; k < len(row); k++ {
			f, err := strconv.ParseFloat(row[k], 64)
			if err != nil {
				return err
			}

			movie.Feature = append(movie.Feature, f)
		}
	}

	return nil
}
