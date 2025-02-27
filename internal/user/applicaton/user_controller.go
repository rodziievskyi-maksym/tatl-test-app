package applicaton

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rodziievskyi-maksym/tatl-test-app/internal/user/repository"
)

type UserController struct {
	userRepository repository.UserRepositoryInterface
}

func (c *UserController) ProfileHandler(ctx *fiber.Ctx) error {
	userProfiles, err := c.userRepository.GetProfiles()
	if err != nil {
		return err
	}

	return ctx.JSON(userProfiles)
}

func NewUserController(userRepository repository.UserRepositoryInterface) UserController {
	return UserController{
		userRepository: userRepository,
	}
}
