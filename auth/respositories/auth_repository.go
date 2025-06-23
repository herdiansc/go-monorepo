package respositories

import (
	"github.com/herdiansc/orderfaz/auth/models"
	"gorm.io/gorm"
)

// AuthRepository struct
type AuthRepository struct {
	db *gorm.DB
}

// NewAuthRepository inits AuthRepository
func NewAuthRepository(db *gorm.DB) AuthRepository {
	return AuthRepository{db: db}
}

// Create saves an auth data
func (repo AuthRepository) Create(auth models.Auth) error {
	result := repo.db.Create(&auth)
	return result.Error
}

// FindByMSISDN finds an auth by msisdn
func (repo AuthRepository) FindByMSISDN(msisdn string) (models.Auth, error) {
	var auth models.Auth
	result := repo.db.Where("msisdn = ?", msisdn).First(&auth)
	return auth, result.Error
}
