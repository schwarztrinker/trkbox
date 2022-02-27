package trk

import (
	"github.com/gofiber/fiber/v2"
	"github.com/schwarztrinker/trkbox/db"
)

func ListAll(c *fiber.Ctx) error {
	user, _ := db.UsersDB.GetUserByUsername(c.Locals("username").(string))

	return c.JSON(user.Timestamps)
}
