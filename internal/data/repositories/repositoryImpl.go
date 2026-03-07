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
