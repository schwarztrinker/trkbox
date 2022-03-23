package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/schwarztrinker/trkbox/conf"
	"github.com/schwarztrinker/trkbox/db"
	"github.com/schwarztrinker/trkbox/router"
)

var ctx = context.Background()

func main() {
	// init fiber
	app := fiber.New()
	//app.Use(cors.New())

	// loading conf from yaml file
	conf.Conf.GetConf()

	// initialize gorm db
	db.InitMariaDB()

	// setup fiber router
	router.SetupRouter(app)

	// run fiber at configured port
	log.Fatal(app.Listen(":" + conf.Conf.Port))
	// nothing will be executed after the fiber app.Listen Command !!!
}
