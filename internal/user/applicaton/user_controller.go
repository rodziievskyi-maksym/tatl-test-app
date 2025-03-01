package applicaton

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rodziievskyi-maksym/tatl-test-app/internal/user/domain/abstractions"
	"github.com/rodziievskyi-maksym/tatl-test-app/internal/user/domain/valueobjects"
)

type UserController struct {
	userRepository abstractions.UserRepositoryInterface
}

func (c *UserController) ProfileHandler(ctx *fiber.Ctx) error {
	usernameValue := ctx.Query("username")

	if usernameValue != "" {
		return c.handleSingleUserProfile(ctx, usernameValue)
	}

	userProfiles, err := c.userRepository.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("failed to get user profiles: %v", err),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(userProfiles)
}

func (c *UserController) handleSingleUserProfile(ctx *fiber.Ctx, username string) error {
	user, err := c.userRepository.GetUserProfileWithFilter(
		valueobjects.UserFilter{Username: username},
	)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("failed to get user profile: %v", err),
		})
	}

	if user == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("user with username %s was not found", username),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

func NewUserController(userRepository abstractions.UserRepositoryInterface) UserController {
	return UserController{
		userRepository: userRepository,
	}
}
