package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v4"
	"github.com/schwarztrinker/trkbox/auth"
	"github.com/schwarztrinker/trkbox/conf"
	"github.com/schwarztrinker/trkbox/trk"

	jwtware "github.com/gofiber/jwt/v3"
)

func SetupRouter(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Your Trkbox instance is running!")
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// API Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Nice to see you using Trkbox! Please Login at /api/auth/login")
	})

	// Auth
	apiAuth := api.Group("/auth")
	apiAuth.Post("/login", auth.Login)
	apiAuth.Post("/create", auth.CreateUser)

	// Trkbox group
	apiTrk := api.Group("/trk")
	// Apply JWT Middleware with signing key
	apiTrk.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(conf.Conf.JwtSecret),
	}))

	apiTrk.Use(func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		name := claims["name"].(string)

		userObj, err := auth.GetUserByUsername(name)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Incorrect User in Token", "data": err})
		}

		c.Locals("user", userObj.Username)
		return c.Next()
	})

	apiTrk.Get("/", auth.Restricted)
	apiTrk.Get("/list/all", trk.ListAll)
	// apiTrk.Get("/list/date")
	// apiTrk.Get("/list/week")

	// apiTrk.Get("/summary/date")
	// apiTrk.Get("/summary/week")

}
