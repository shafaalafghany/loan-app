package util

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/shafaalafghany/loan-app/model"
)

func GetUserIdFromToken(c *fiber.Ctx) (string, error) {
	userId := c.Locals("userId")
	if userId == nil {
		return "", fiber.NewError(fiber.StatusUnauthorized, "id not found in token")
	}

	userIdStr, ok := userId.(string)
	if !ok {
		return "", fiber.NewError(fiber.StatusUnauthorized, "invalid id format")
	}

	return userIdStr, nil
}

func GetDefaultModelAuditLog(action, resource string, user model.User) *model.AuditLog {
	return &model.AuditLog{
		ID:        uuid.New(),
		UserID:    user.ID,
		Action:    action,
		Resource:  resource,
		Timestamp: time.Now(),
	}
}
