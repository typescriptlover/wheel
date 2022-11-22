package server

import (
	"wheel/controllers"

	"github.com/gofiber/fiber/v2"
)

func CreateRouter() *fiber.App {
	r := fiber.New()

	controllers.AuthGroup(r)

	return r
}
