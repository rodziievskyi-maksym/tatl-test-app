package fiber_server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type FiberHTTPServer struct {
	app     *fiber.App
	address string
}

func NewApp() *fiber.App {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(recover.New())

	return app
}

func NewFiberHTTPServer(app *fiber.App, address string) FiberHTTPServer {

	return FiberHTTPServer{
		app:     app,
		address: address,
	}
}

func (s FiberHTTPServer) Start() error {
	if err := s.app.Listen(s.address); err != nil {
		return fmt.Errorf("failed to start HTTP server: %s", err.Error())
	}

	return nil
}
