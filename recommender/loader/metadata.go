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

// LoadMetadata grabs IMDB and TMDB IDs from CSV and load them into movies in memory.
func LoadMetadata() error {
	csvFile, err := os.Open(fmt.Sprintf("%s/links.csv", dataDir))
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

		id, err := strconv.ParseInt(row[0], 10, 64)
		if err != nil {
			continue
		}

		if _, ok := movies[model.MovieID(id)]; !ok {
			continue
		}

		movies[model.MovieID(id)].IMDBID = "tt" + row[1]
		movies[model.MovieID(id)].TMDBID = row[2]
	}

	return nil
}

// LoadTags grabs tags from CSV and load them into movies.
func LoadTags() error {
	csvFile, err := os.Open(fmt.Sprintf("%s/tags.csv", dataDir))
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

		id, err := strconv.ParseInt(row[1], 10, 64)
		if err != nil {
			continue
		}

		if _, ok := movies[model.MovieID(id)]; !ok {
			continue
		}

		if movies[model.MovieID(id)].Tags == nil {
			movies[model.MovieID(id)].Tags = []string{}
		}

		movies[model.MovieID(id)].Tags = append(movies[model.MovieID(id)].Tags, row[2])
	}

	return nil
}
