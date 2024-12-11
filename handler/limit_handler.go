package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shafaalafghany/loan-app/service"
	"go.uber.org/zap"
)

type LimitHandler struct {
	s   service.LimitServiceInterface
	log *zap.Logger
}

func NewLimitHandler(s service.LimitServiceInterface, log *zap.Logger) *LimitHandler {
	return &LimitHandler{
		s:   s,
		log: log,
	}
}

func (h *LimitHandler) Get(c *fiber.Ctx) error {
	h.log.Info("incoming request in handler for getting user limit")
	return h.s.Get(c)
}
