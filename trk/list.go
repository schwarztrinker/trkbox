package trk

import (
	"github.com/gofiber/fiber/v2"
	"github.com/schwarztrinker/trkbox/db"
)

func ListAll(c *fiber.Ctx) error {
	user, _ := db.UsersDB.GetUserByUsername(c.Locals("username").(string))
	return c.JSON(user.Timestamps)
}

func ListDate(c *fiber.Ctx) error {
	user, _ := db.UsersDB.GetUserByUsername(c.Locals("username").(string))
	date := c.Params("date")

	ts, err := user.Timestamps.GetTimestampsByDay(date)
	if err != nil {
		return c.JSON(fiber.Map{"status": "error", "message": "No timestamps found for this date", "data": nil})
	}

	return c.JSON(ts)
}
