package models

import "time"

type Todo struct {
	Id        string `gorm:"primary_key"`
	Content   string
	CreatedAt time.Time
	CreatedBy string
	User      User `gorm:"foreignKey:CreatedBy"`
}
