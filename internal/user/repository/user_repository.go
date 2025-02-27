package repository

import (
	"fmt"
	"github.com/rodziievskyi-maksym/tatl-test-app/internal/user/domain"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	GetProfiles() ([]domain.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) GetProfiles() ([]domain.User, error) {
	var users []domain.User

	result := r.db.
		Preload("Profile").
		Preload("Data").
		Find(&users)

	if result.Error != nil {
		return nil, fmt.Errorf("error occurs while retrieving users from db: %s", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return users, nil
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{
		db: db,
	}
}
