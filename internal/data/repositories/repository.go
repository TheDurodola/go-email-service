package repositories

import "github.com/TheDurodola/go-email-service/internal/data/models"

type EmailRepository interface {
    CreateTemplate(template *models.EmailTemplate) error
	CreateLog(log *models.OutgoingEmail) error
}