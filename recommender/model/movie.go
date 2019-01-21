package model

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// AllMovies fetches all movies from database.
func AllMovies() ([]*Movie, error) {
	all := []*Movie{}

	if err := db.Find(&all).Error; err != nil {
		return nil, err
	}

	return all, nil
}

// Movie is a model for movie entity.
type Movie struct {
	gorm.Model
	TMDBID string `gorm:"column:tmdb_id"`
	IMDBID string `gorm:"column:imdb_id"`

	Title string `gorm:"column:title"`
	Year  int    `gorm:"column:year"`

	NumRating     int32   `gorm:"column:num_rating"`
	IMDBRating    float32 `gorm:"column:imdb_rating"`
	AverageRating float32 `gorm:"column:average_rating"`

	// Feature is k-dimensional vector that represents the latent feature of a movie.
	Feature pq.Float64Array `gorm:"column:feature"`
	Tags    pq.StringArray  `gorm:"column:tags"`

	Cluster          int64         `gorm:"column:cluster"`
	NearestClusters  pq.Int64Array `gorm:"column:nearest_clusters"`
	FarthestClusters pq.Int64Array `gorm:"column:farthest_clusters"`

	Ratings []*Rating `gorm:"foreignkey:MovieID"`
}

// TableName returns database table name which this entity is mapping to.
func (*Movie) TableName() string {
	return "movies"
}

// MovieDetail contains poster url and high level description of a movie.
type MovieDetail struct {
	IMDBID    string          `gorm:"column:imdb_id; primary_key"`
	CreatedAt time.Time       `gorm:"column:created_at"`
	UpdatedAt time.Time       `gorm:"column:updated_at"`
	Detail    json.RawMessage `gorm:"column:detail"`
}

// TableName returns database table name which this entity is mapping to.
func (*MovieDetail) TableName() string {
	return "movie_details"
}

// MovieTrailer contains video URL of a movie trailer.
type MovieTrailer struct {
	IMDBID    string          `gorm:"column:imdb_id; primary_key"`
	CreatedAt time.Time       `gorm:"column:created_at"`
	UpdatedAt time.Time       `gorm:"column:updated_at"`
	Trailer   json.RawMessage `gorm:"column:trailer"`
}

// TableName returns database table name which this entity is mapping to.
func (*MovieTrailer) TableName() string {
	return "movie_trailers"
}
