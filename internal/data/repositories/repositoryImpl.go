package repositories

import (
	"github.com/TheDurodola/go-email-service/internal/data/models"
	"gorm.io/gorm"
)

type gormEmailRepo struct {
    db *gorm.DB
}

func NewEmailRepository(db *gorm.DB) *gormEmailRepo {
    return &gormEmailRepo{db: db}
}

func (r *gormEmailRepo) CreateLog(log *models.OutgoingEmail) error {
    return r.db.Create(log).Error
}

func (r *gormEmailRepo) CreateTemplate(template *models.EmailTemplate) error {
    return r.db.Create(template).Error
}