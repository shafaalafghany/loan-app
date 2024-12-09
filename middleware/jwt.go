package middleware

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware(key string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "missing authorization header",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid authorization header format",
			})
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "unexpected signin method")
			}
			return []byte(key), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid or expired token",
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Locals("userId", claims["id"])
			if exp, ok := claims["exp"].(float64); ok {
				if time.Unix(int64(exp), 0).Before(time.Now()) {
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
						"message": "token has expired",
					})
				}
			} else {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "invalid token claims",
				})
			}
		}

		return c.Next()
	}
}
