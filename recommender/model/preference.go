package model

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// FetchUserPreference fetches user preference from database.
func FetchUserPreference(email string) (*Preference, error) {
	pref := &Preference{}

	if err := db.Where("user_email = ?", email).Find(pref).Error; err != nil {
		return nil, err
	}

	return pref, nil
}

// Preference is a user preference for movie.
type Preference struct {
	gorm.Model
	UserEmail string          `gorm:"column:user_email"`
	Value     pq.Float64Array `gorm:"column:value"`
}

// TableName returns database table name which this entity is mapping to.
func (*Preference) TableName() string {
	return "preferences"
}
