package repositories

import (
	"github.com/TheDurodola/go-email-service/internal/data/models"
	"gorm.io/gorm"
)

type GormEmailRepo struct {
	db *gorm.DB
}

func NewEmailRepository(db *gorm.DB) *GormEmailRepo {
	return &GormEmailRepo{db: db}
}

func (r *GormEmailRepo) CreateLog(log *models.OutgoingEmail) error {
	return r.db.Create(log).Error
}

func (r *GormEmailRepo) CreateTemplate(template *models.EmailTemplate) error {
	return r.db.Create(template).Error
}


func (r *GormEmailRepo) GetTemplateByName(name string) (*models.EmailTemplate, error) {
	var template models.EmailTemplate
	if err := r.db.Where("template_name = ?", name).First(&template).Error; err != nil {
		return nil, err
	}
	return &template, nil
}