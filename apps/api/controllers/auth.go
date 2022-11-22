package controllers

import "github.com/gofiber/fiber/v2"

func AuthGroup(r *fiber.App) {
	auth := r.Group("/auth")

	auth.Get("/", Index)
}

func Index(c *fiber.Ctx) error {
	return c.SendString("yo")
}
