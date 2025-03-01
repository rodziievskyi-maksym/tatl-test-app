package abstractions

import "github.com/rodziievskyi-maksym/tatl-test-app/internal/auth/domain/entities"

type AuthRepositoryContract interface {
	FindByAPIKey(apiKey string) (*entities.Auth, error)
}
