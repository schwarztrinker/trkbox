package trk

import (
	"github.com/gofiber/fiber/v2"
	"github.com/schwarztrinker/trkbox/db"
)

func ListAll(c *fiber.Ctx) error {
	user, _ := db.GetUserByUsername(c.Locals("username").(string))
	ts := db.GetTimestampsFromUser(*user)

	return c.JSON(fiber.Map{"status": "success", "message": "Timestamps available for this date", "data": ts})
}

func ListDate(c *fiber.Ctx) error {
	user, _ := db.GetUserByUsername(c.Locals("username").(string))
	date := c.Params("date")

	ts, err := db.GetTimestampsByDay(*user, date)
	if err != nil {
		return c.JSON(fiber.Map{"status": "error", "message": "No timestamps found for this date", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Timestamps for this date", "data": ts})
}

func ListWeek(c *fiber.Ctx) error {
	user, _ := db.GetUserByUsername(c.Locals("username").(string))
	week := c.Params("week")

	ts, err := db.GetTimestampsByWeek(*user, week)
	if err != nil {
		return c.JSON(fiber.Map{"status": "error", "message": "No timestamps found for this date", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Timestamps for this date", "data": ts})
}
