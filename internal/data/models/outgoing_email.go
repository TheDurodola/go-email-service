package models

import "gorm.io/gorm"

type OutgoingEmail struct {
	gorm.Model
	TemplateName string `gorm:"unique;not null"`
	Subject      string `gorm:"not null"`
	Body         string `gorm:"not null"`
}