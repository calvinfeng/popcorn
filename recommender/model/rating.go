package model

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Rating struct {
	gorm.Model
	MovieID   uint    `gorm:"column:movie_id"`
	UserEmail string  `gorm:"column:user_email"`
	Value     float32 `gorm:"column:value"`
}

func (*Rating) TableName() string {
	return "ratings"
}

type Preference struct {
	gorm.Model
	UserEmail uint            `gorm:"column:user_email"`
	Value     pq.Float64Array `gorm:"column:value"`
}

func (*Preference) TableName() string {
	return "preferences"
}
