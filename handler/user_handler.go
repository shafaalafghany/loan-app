package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shafaalafghany/loan-app/model"
	"github.com/shafaalafghany/loan-app/service"
	"go.uber.org/zap"
)

type UserHandler struct {
	s   service.UserServiceInterface
	log *zap.Logger
}

func NewUserHandler(s service.UserServiceInterface, log *zap.Logger) *UserHandler {
	return &UserHandler{
		s:   s,
		log: log,
	}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	h.log.Info("incoming request in handler for register new user")
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to parsing payload",
		})
	}

	if err := h.s.Register(c, &user); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "create new user successfully",
	})
}
