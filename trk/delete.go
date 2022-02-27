package trk

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/schwarztrinker/trkbox/db"
)

func DeleteTimestamp(c *fiber.Ctx) error {
	user, _ := db.UsersDB.GetUserByUsername(c.Locals("username").(string))
	id := c.Params("uuid")

	parsedUid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(fiber.Map{"status": "error", "message": "Timestamp Id could not be parsed", "data": nil})
	}

	user.Timestamps.DeleteTimestampByUuid(parsedUid)

	return c.JSON(fiber.Map{"status": "success", "message": "Timestamp deleted", "data": parsedUid})
}
