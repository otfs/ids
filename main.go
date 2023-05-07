package main

import (
	"ids/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.InitConfig()
	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())
	initRoute(app)
	log.Fatal(app.Listen(config.ServerListenAddr))
}
