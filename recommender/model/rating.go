package model

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// Rating is a movie rating submitted by user.
type Rating struct {
	gorm.Model
	MovieID   uint    `gorm:"column:movie_id"`
	UserEmail string  `gorm:"column:user_email"`
	Value     float32 `gorm:"column:value"`
}

// TableName returns database table name which this entity is mapping to.
func (*Rating) TableName() string {
	return "ratings"
}

// Preference is a user preference for movie.
type Preference struct {
	gorm.Model
	UserEmail uint            `gorm:"column:user_email"`
	Value     pq.Float64Array `gorm:"column:value"`
}

// TableName returns database table name which this entity is mapping to.
func (*Preference) TableName() string {
	return "preferences"
}
