package api

import (
	"github.com/rodziievskyi-maksym/tatl-test-app/internal/user/applicaton"
	"github.com/rodziievskyi-maksym/tatl-test-app/pkg/infrastructure/server/fiber_server"
)

type HTTPServer interface {
	Start() error
}

func NewHTTPServer(address string, userController applicaton.UserController) HTTPServer {
	app := fiber_server.NewApp()

	app.Get("/profiles", userController.ProfileHandler)

	return fiber_server.NewFiberHTTPServer(app, address)
}
