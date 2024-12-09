package util

import "github.com/gofiber/fiber/v2"

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
