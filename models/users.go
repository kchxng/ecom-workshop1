package models

import "time"

type Users struct {
	Id        uint       `gorm:"primaryKey; autoIncrement:true" json:"id"`
	Username  string     `gorm:"size:255;not null" json:"username"`
	Password  string     `gorm:"size:255;not null" json:"password"`
	Email     *string    `gorm:"size:255;null" json:"email"`
	CreatedAt *time.Time `gorm:"null" json:"created_at"`
	UpdatedAt *time.Time `gorm:"null" json:"updated_at"`
	DeletedAt *time.Time `gorm:"null" json:"deleted_at"`
}
