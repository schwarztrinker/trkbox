package handler

import "github.com/gofiber/fiber/v2"

func Root(c *fiber.Ctx) error {
	return c.SendString("Your Trkbox instance is running!")
}

func Pong(c *fiber.Ctx) error {
	return c.SendString("pong")
}

func Api(c *fiber.Ctx) error {
	return c.SendString("Nice to see you using Trkbox! Please Login at /api/auth/login")
}
