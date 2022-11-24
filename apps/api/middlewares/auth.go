package middlewares

import (
	"fmt"
	"strings"
	"wheel/config"
	"wheel/db"
	"wheel/errors"
	"wheel/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AuthGuard(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")

	if len(authorization) <= 0 {
		fmt.Println("No authorization")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"type":    errors.AuthError,
			"message": "Not logged in",
		})
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) != 2 {
		fmt.Println("not 2")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"type":    errors.AuthError,
			"message": "Not logged in",
		})
	}

	claims := jwt.MapClaims{}
	if _, err := jwt.ParseWithClaims(token[1], claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().JWT_SECRET), nil
	}); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"type":    errors.AuthError,
			"message": "Not logged in",
		})
	}

	user := models.User{}
	if result := db.GetDB().Select("id", "username", "email", "updated", "created").Where(&models.User{
		ID: claims["id"].(string),
	}).First(&user); result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"type":    errors.AuthError,
			"message": "Not logged in",
		})
	}

	c.Locals("user", user)

	return c.Next()
}
