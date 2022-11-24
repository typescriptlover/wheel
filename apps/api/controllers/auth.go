package controllers

import (
	"crypto/rand"
	"fmt"
	"wheel/config"
	"wheel/db"
	"wheel/errors"
	"wheel/middlewares"
	"wheel/models"
	"wheel/schema"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lucsky/cuid"
	"golang.org/x/crypto/bcrypt"
)

// Auth group
func AuthGroup(r *fiber.App) {
	auth := r.Group("/auth")

	auth.Post("/register", Register)
	auth.Post("/login", Login)

	me := auth.Group("/me", middlewares.AuthGuard)
	me.Get("/", Me)
}

// Register route
func Register(c *fiber.Ctx) error {
	body := new(schema.Register)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"type":    errors.ParseError,
			"message": err.Error(),
		})
	}

	if err := schema.ValidateSchema(*body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"type":    errors.ValidationError,
			"message": err,
		})
	}

	id, err := cuid.NewCrypto(rand.Reader)
	if err != nil {
		fmt.Printf("%v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"type":    errors.InternalError,
			"message": "Failed generating ID",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("%v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"type":    errors.InternalError,
			"message": "Failed hashing password",
		})
	}

	user := &models.User{
		ID:       id,
		Username: body.Username,
		Email:    body.Email,
		Password: string(hashedPassword),
	}

	if result := db.GetDB().Create(user); result.Error != nil {
		fmt.Printf("%v\n", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"type":    errors.DBError,
			"message": "Failed creating user",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})

	tokenString, err := token.SignedString([]byte(config.GetConfig().JWT_SECRET))
	if err != nil {
		fmt.Printf("%v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"type":    errors.InternalError,
			"message": "Failed signing token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "user created",
		"payload": fiber.Map{
			"token": tokenString,
		},
	})
}

// Login route
func Login(c *fiber.Ctx) error {
	body := new(schema.Login)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"type":    errors.ParseError,
			"message": err.Error(),
		})
	}

	if err := schema.ValidateSchema(*body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"type":    errors.ValidationError,
			"message": "Invalid username or password",
		})
	}

	user := models.User{}

	if result := db.GetDB().Where(&models.User{
		Username: body.Username,
	}).First(&user); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"type":    errors.ValidationError,
			"message": "Invalid username or password",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"type":    errors.ValidationError,
			"message": "Invalid username or password",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})

	tokenString, err := token.SignedString([]byte(config.GetConfig().JWT_SECRET))
	if err != nil {
		fmt.Printf("%v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"type":    errors.InternalError,
			"message": "Failed signing token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "logged in",
		"payload": fiber.Map{
			"token": tokenString,
		},
	})
}

func Me(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"payload": fiber.Map{
			"user": c.Locals("user"),
		},
	})
}
