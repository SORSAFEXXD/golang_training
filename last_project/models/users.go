package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserProfile struct {
	gorm.Model
	EmployeeID uuid.UUID `gorm:"type:char(36);uniqueIndex;not null" json:"employee_id"`
	Name       string    `gorm:"not null" json:"name"`
	Lastname   string    `gorm:"not null" json:"lastname"`
	Birthday   string    `gorm:"type:date;not null" json:"birthday"`
	Age        int       `gorm:"not null" json:"age"`
	Email      string    `gorm:"not null" json:"email"`
	Tel        string    `gorm:"not null" json:"tel"`
}

type ProfileInput struct {
	EmployeeID uuid.UUID `json:"employee_id"`
	Name       string    `json:"name"`
	Lastname   string    `json:"lastname"`
	Birthday   string    `json:"birthday"`
	Email      string    `json:"email"`
	Tel        string    `json:"tel"`
}
