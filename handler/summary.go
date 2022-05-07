package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/schwarztrinker/trkbox/db"
	"github.com/schwarztrinker/trkbox/summary"
)

func GetSummaryByDate(c *fiber.Ctx) error {
	date := c.Params("date")
	user, err := db.GetUserByUsername(c.Locals("username").(string))
	if err != nil {
		return c.JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}

	summary, err := summary.GenerateSummaryByDate(user, date)
	if err != nil {
		return c.JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Summary for requested date", "data": summary})

}

func GetSummaryByWeek(c *fiber.Ctx) error {
	week := c.Params("week")
	user, err := db.GetUserByUsername(c.Locals("username").(string))
	if err != nil {
		return c.JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}

	summary, err := summary.GenerateSummaryByWeek(user, week)
	if err != nil {
		return c.JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Summary for requested date", "data": summary})

}
