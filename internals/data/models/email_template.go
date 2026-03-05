package models

import "gorm.io/gorm"

type EmailTemplate struct {
	gorm.Model
	Name    string `gorm:"unique;not null"`
	Subject string `gorm:"not null"`
	Body    string `gorm:"not null"`
}