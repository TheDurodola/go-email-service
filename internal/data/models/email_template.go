package models

import "gorm.io/gorm"


type TemplateType string

const (
    Static   TemplateType = "STATIC"
    Dynamic   TemplateType = "DYNAMIC"
)
type EmailTemplate struct {
	gorm.Model
	Name    string `gorm:"unique;not null"`
	Type   TemplateType `gorm:"not null"`
	Subject string `gorm:"not null"`
	Body    string `gorm:"not null"`
}