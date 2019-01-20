package loader

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"popcorn/recommender/model"
	"regexp"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
)

// LoadMovies grabs movie data from CSV file into memory.
func LoadMovies() error {
	csvFile, err := os.Open(fmt.Sprintf("%s/movies.csv", dataDir))
	if err != nil {
		return err
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))
	yearPattern, _ := regexp.Compile("\\(\\d{4}\\)")
	numericPattern, _ := regexp.Compile("\\d{4}")

	for {
		row, err := reader.Read()
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		var id, year int64
		id, err = strconv.ParseInt(row[0], 10, 64)
		if err != nil {
			continue
		}

		yearStr := yearPattern.FindString(row[1])

		year, err = strconv.ParseInt(numericPattern.FindString(yearStr), 10, 64)
		if err != nil {
			continue
		}

		movies[MovieID(id)] = &model.Movie{
			Model: gorm.Model{ID: uint(id)},
			Year:  int(year),
			Title: strings.Trim(row[1], " "+yearStr),
		}
	}

	return nil
}
