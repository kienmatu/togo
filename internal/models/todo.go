package models

import "time"

type Location struct {
	Id        string `gorm:"primary_key"`
	Name      string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy string
	Lat       float64 `gorm:"type:decimal(10,8)"`
	Lng       float64 `gorm:"type:decimal(11,8)"`
	User      User    `gorm:"foreignKey:CreatedBy"`
	Note      string
}
