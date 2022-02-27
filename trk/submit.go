package trk

import (
	"github.com/gofiber/fiber/v2"
	"github.com/schwarztrinker/trkbox/auth"
	"github.com/schwarztrinker/trkbox/db"
)

func SubmitTimestamp(c *fiber.Ctx) error {
	user, err := auth.GetUserByUsername(c.Locals("username").(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "User could not be found"})
	}

	var t db.Timestamp

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	if err := c.BodyParser(&t); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Timestamp could not be parsed"})
	}

	timestamps := user.Timestamps

	timestamps.Timestamps = append(timestamps.Timestamps, t)

	return c.JSON(fiber.Map{"status": "success", "message": "Timestamp saved successful", "data": t})
}
