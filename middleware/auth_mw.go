package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shanomz7235/bookstore-back/utils"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing token",
		})
	}

	token, err := utils.ValidateJWT(cookie)

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token",
		})
	}
	
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid claims",
		})
	}


	c.Locals("user_email", claims["user_email"])
	c.Locals("role", claims["user_role"])
	c.Locals("claims", claims)


	return c.Next()
}

func RoleRequired(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("role").(string)
		for _, r := range roles {
			if r == role {
				return c.Next()
			}
		}
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "access denied",
		})
	}
}

