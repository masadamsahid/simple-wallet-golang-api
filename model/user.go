package model

import "time"

type User struct {
	ID        uint      `gorm:"column:id;primary_key"`
	Name      string    `gorm:"type:text;not null"`
	Email     string    `gorm:"type:text;uniqueIndex;not null"`
	Password  string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (User) TableName() string {
	return "users"
}
