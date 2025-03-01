package api

import (
	"github.com/rodziievskyi-maksym/tatl-test-app/api/middlewares"
	"github.com/rodziievskyi-maksym/tatl-test-app/internal/auth/domain/abstractions"
	"github.com/rodziievskyi-maksym/tatl-test-app/internal/user/applicaton"
	"github.com/rodziievskyi-maksym/tatl-test-app/pkg/infrastructure/server/fiber_server"
)

type HTTPServer interface {
	Start() error
}

func NewHTTPServer(
	address string,
	userController applicaton.UserController,
	authRepository abstractions.AuthRepositoryContract,
) HTTPServer {
	app := fiber_server.NewApp()
	app.Use(middlewares.AuthMiddleware(authRepository))

	app.Get("/profiles", userController.ProfileHandler)

	return fiber_server.NewHTTPServer(app, address)
}
