package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/schwarztrinker/trkbox/conf"
	"github.com/schwarztrinker/trkbox/db"
	"github.com/schwarztrinker/trkbox/router"
)

func main() {
	// init fiber
	app := fiber.New()
	//app.Use(cors.New())

	// loading conf from yaml file
	conf.Conf.GetConf()

	// Loading Timestamp DB
	db.ConnectDB()
	//db.LoadingTimestampsGlobalFromDB()
	db.UsersDB.LoadUserDB()

	// setup fiber router
	router.SetupRouter(app)

	// run fiber at configured port
	log.Fatal(app.Listen(":" + conf.Conf.Port))
}
