package handler

import "github.com/gofiber/fiber/v2"

func Root(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Your Trkbox instace ist running", "data": nil})
}

func Pong(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Pong", "data": nil})
}

func Api(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Nice to see you using Trkbox! Please Login at /api/auth/login", "data": nil})
}
