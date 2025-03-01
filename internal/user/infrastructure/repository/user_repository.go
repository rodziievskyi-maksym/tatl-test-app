package repository

import (
	"fmt"
	"github.com/rodziievskyi-maksym/tatl-test-app/internal/user/domain/abstractions"
	"github.com/rodziievskyi-maksym/tatl-test-app/internal/user/domain/aggregates"
	"github.com/rodziievskyi-maksym/tatl-test-app/internal/user/domain/valueobjects"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) GetUserProfileWithFilter(filter valueobjects.UserFilter) (*aggregates.User, error) {
	var user *aggregates.User

	result := r.db.
		Preload("Profile").
		Preload("Data")

	if filter.Username != "" {
		result.Where("username = ?", filter.Username).Find(&user)
	}

	if result.Error != nil {
		return nil, fmt.Errorf("error occurs while retrieving user from db: %s", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return user, nil
}

func (r *UserRepository) GetAll() ([]aggregates.User, error) {
	var users []aggregates.User

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

func NewUserRepository(db *gorm.DB) abstractions.UserRepositoryInterface {
	return &UserRepository{
		db: db,
	}
}
