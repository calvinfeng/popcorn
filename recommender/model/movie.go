package model

import (
	"encoding/json"
	"time"

	"github.com/lib/pq"
)

// FetchAllMovies fetches all movies from database.
func FetchAllMovies() (movies []*Movie, err error) {
	err = db.Find(&movies).Error
	if err != nil {
		return
	}

	return
}

// FetchMovies fetches a list of movies by ID(s).
func FetchMovies(ids []MovieID) (movies []*Movie, err error) {
	err = db.Where("id in (?)", ids).Find(&movies).Error
	if err != nil {
		return
	}

	return
}

type (
	// MovieID is an alias to avoid confusing uint with UserID.
	MovieID uint

	// UserID is an alias to avoid confusing uint with MovieID.
	UserID uint
)

// Movie is a model for movie entity.
type Movie struct {
	ID        MovieID    `gorm:"column:id;primary_key"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
	TMDBID    string     `gorm:"column:tmdb_id"`
	IMDBID    string     `gorm:"column:imdb_id"`

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
