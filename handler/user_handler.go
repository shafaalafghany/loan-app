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
	var user model.UserRequest
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to parsing payload",
		})
	}

	return h.s.Register(c, &user)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	h.log.Info("incoming request in handler for user login")
	var req model.UserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to parsing payload",
		})
	}

	return h.s.Login(c, &req)
}
