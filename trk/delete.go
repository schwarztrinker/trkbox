package trk

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/schwarztrinker/trkbox/db"
)

func DeleteTimestamp(c *fiber.Ctx) error {
	//user, _ := db.GetUserByUsername(c.Locals("username").(string))
	id := c.Params("uuid")
	user, _ := db.GetUserByUsername(c.Locals("username").(string))

	parsedUid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(fiber.Map{"status": "error", "message": "Timestamp Id could not be parsed", "data": nil})
	}

	_, err = db.DeleteTimestampByUuid(*user, parsedUid)
	if err != nil {
		return c.JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Timestamp deleted", "data": parsedUid})
}
