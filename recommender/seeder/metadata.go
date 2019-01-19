package seeder

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// LoadMetadata grabs IMDB and TMDB IDs from CSV and load them into movies.
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

		if _, ok := movies[uint(id)]; !ok {
			continue
		}

		movies[uint(id)].IMDBID = "tt" + row[1]
		movies[uint(id)].TMDBID = row[2]
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

		if _, ok := movies[uint(id)]; !ok {
			continue
		}

		if movies[uint(id)].Tags == nil {
			movies[uint(id)].Tags = []string{}
		}

		movies[uint(id)].Tags = append(movies[uint(id)].Tags, row[2])
	}

	return nil
}
