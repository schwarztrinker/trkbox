package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/schwarztrinker/trkbox/auth"
	"github.com/schwarztrinker/trkbox/conf"
	"github.com/schwarztrinker/trkbox/handler"
	"github.com/schwarztrinker/trkbox/trk"

	jwtware "github.com/gofiber/jwt/v3"
)

func SetupRouter(app *fiber.App) {
	app.Get("/", handler.Root)
	app.Get("/ping", handler.Pong)

	// /api Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Api)

	// /api/auth
	apiAuth := api.Group("/auth")
	apiAuth.Post("/login", auth.Login)
	apiAuth.Post("/create", auth.CreateUser)

	// /api/trk
	apiTrk := api.Group("/trk")
	// Apply JWT Middleware with signing key to all /api/trk requests
	apiTrk.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(conf.Conf.JwtSecret),
	}))

	// Middleware to check if User to the provided token exists
	// forces check on /api/trk/*
	apiTrk.Use("/", auth.CheckUserFromToken)

	// /api/trk application routes
	apiTrk.Get("/", auth.Restricted)
	apiTrk.Get("/whoami", auth.Whoami)
	apiTrk.Get("/list/all", trk.ListAll)
	apiTrk.Get("/list/date/:date", trk.ListDate)

	apiTrk.Post("/submit", trk.SubmitTimestamp)
	apiTrk.Post("/delete/:uuid", trk.DeleteTimestamp)

	// apiTrk.Get("/list/week")

	// apiTrk.Get("/summary/date")
	// apiTrk.Get("/summary/week")

}
