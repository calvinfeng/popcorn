package model

import (
	"github.com/jinzhu/gorm"
)

// FetchUserRatings fetches a list of ratings by a user using his/her email.
func FetchUserRatings(email string) ([]*Rating, error) {
	ratings := []*Rating{}

	if err := db.Find(&ratings).Where("user_email = ?", email).Error; err != nil {
		return nil, err
	}

	return ratings, nil
}

// InsertUserRating inserts a movie rating.
func InsertUserRating(email string, id MovieID, val float32) error {
	rating := &Rating{
		UserEmail: email,
		MovieID:   id,
		Value:     val,
	}

	return db.Create(rating).Error
}

// Rating is a movie rating submitted by user.
type Rating struct {
	gorm.Model
	UserEmail string  `gorm:"column:user_email"`
	MovieID   MovieID `gorm:"column:movie_id"`
	Value     float32 `gorm:"column:value"`
}

// TableName returns database table name which this entity is mapping to.
func (*Rating) TableName() string {
	return "ratings"
}
