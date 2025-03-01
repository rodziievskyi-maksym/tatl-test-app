package abstractions

import (
	"github.com/rodziievskyi-maksym/tatl-test-app/internal/user/domain/aggregates"
	"github.com/rodziievskyi-maksym/tatl-test-app/internal/user/domain/valueobjects"
)

type UserRepositoryInterface interface {
	GetAll() ([]aggregates.User, error)
	GetUserProfileWithFilter(filter valueobjects.UserFilter) (*aggregates.User, error)
}
