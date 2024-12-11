package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shafaalafghany/loan-app/repository"
	"github.com/shafaalafghany/loan-app/util"
	"go.uber.org/zap"
)

type LimitServiceInterface interface {
	Get(*fiber.Ctx) error
}

type LimitService struct {
	repo repository.LimitRepositoryInterface
	log  *zap.Logger
}

func NewLimitService(repo repository.LimitRepositoryInterface, log *zap.Logger) LimitServiceInterface {
	return &LimitService{
		repo: repo,
		log:  log,
	}
}

func (s *LimitService) Get(c *fiber.Ctx) error {
	s.log.Info("incoming request to get user limit")
	userId, err := util.GetUserIdFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	limits, err := s.repo.GetByUserId(userId)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to get user limit",
		})
	}

	return c.Status(fiber.StatusOK).JSON(limits)
}
