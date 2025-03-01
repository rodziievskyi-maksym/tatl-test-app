package repository

import (
	"fmt"
	"github.com/rodziievskyi-maksym/tatl-test-app/internal/auth/domain/abstractions"
	"github.com/rodziievskyi-maksym/tatl-test-app/internal/auth/domain/entities"
	"gorm.io/gorm"
)

var _ abstractions.AuthRepositoryContract = (*AuthRepository)(nil)

type AuthRepository struct {
	db *gorm.DB
}

func (r *AuthRepository) FindByAPIKey(apiKey string) (*entities.Auth, error) {
	auth := new(entities.Auth)

	result := r.db.Where("`api-key` = ?", apiKey).Find(auth)

	if result.Error != nil {
		return nil, fmt.Errorf("error occurs while retrieving auth from db: %s", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return auth, nil
}

func NewAuthRepository(db *gorm.DB) abstractions.AuthRepositoryContract {
	return &AuthRepository{
		db: db,
	}
}
