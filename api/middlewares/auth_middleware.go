package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rodziievskyi-maksym/tatl-test-app/internal/auth/domain/abstractions"
)

func AuthMiddleware(authRepository abstractions.AuthRepositoryContract) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		apiKey := ctx.Get("Api-key")
		if apiKey == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "missing Api-key value",
			})
		}

		auth, err := authRepository.FindByAPIKey(apiKey)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("failed to get auth data: %v", err),
			})
		}

		if auth == nil {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "access forbidden",
			})
		}

		return ctx.Next()
	}
}
