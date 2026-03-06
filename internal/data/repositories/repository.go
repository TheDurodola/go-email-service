package repositories

import "github.com/TheDurodola/go-email-service/internal/data/models"

type EmailRepository interface {
    CreateEmailTemplate(template *models.EmailTemplate) error
	GetEmailTemplateByName(name string) (*models.EmailTemplate, error)
	CreateEmailLog(log *models.OutgoingEmail) error
	GetEmailLogByID(id uint)(*models.OutgoingEmail, error)
}