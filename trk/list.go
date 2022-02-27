package trk

import (
	"github.com/gofiber/fiber/v2"
	"github.com/schwarztrinker/trkbox/auth"
)

func ListAll(c *fiber.Ctx) error {
	user, _ := auth.GetUserByUsername(c.Locals("username").(string))

	return c.JSON(user.Timestamps)
}
