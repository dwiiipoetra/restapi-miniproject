package middleware

import "github.com/gofiber/fiber/v2"

func Auth(context *fiber.Ctx) error {
	token := context.Get("x-token")
	if token != "secret" {
		return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized user"})
	}
	return context.Next()
}

func PermissionCreate(context *fiber.Ctx) error {
	return context.Next()
}
