package models

import "gorm.io/gorm"

type OutgoingEmail struct {
	gorm.Model
	TemplateName string `gorm:"unique;not null"`
	AppName string `gorm:"not null"`
	Recipient  string `gorm:"not null"`
	FirstName  string `gorm:"not null"`
}