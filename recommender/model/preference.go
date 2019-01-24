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

// InsertUpdateUserPreference will either create or update a user preference.
func InsertUpdateUserPreference(email string, val []float64) error {
	pref, err := FetchUserPreference(email)
	if err == nil {
		return db.Model(pref).Update("value", val).Error
	}

	pref = &Preference{
		UserEmail: email,
		Value:     val,
	}

	return db.Create(pref).Error
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
