package respositories

import (
	"github.com/herdiansc/orderfaz/auth/models"
	"gorm.io/gorm"
)

// LogisticRepository struct
type LogisticRepository struct {
	db *gorm.DB
}

// NewLogisticRepository inits LogisticRepository
func NewLogisticRepository(db *gorm.DB) LogisticRepository {
	return LogisticRepository{db: db}
}

// List finds list of all logistics by filter
func (repo LogisticRepository) List(filter map[string]interface{}) ([]models.Logistic, error) {
	var data []models.Logistic
	result := repo.db.Where(filter).Order("amount desc").Find(&data)
	return data, result.Error
}

// FindByUUID finds a logistic by msisdn
func (repo LogisticRepository) FindByUUID(uuid string) (models.Logistic, error) {
	var data models.Logistic
	result := repo.db.Where("uuid = ?", uuid).First(&data)
	return data, result.Error
}
