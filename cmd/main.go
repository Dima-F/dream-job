package main

import (
	"github.com/Dima-F/dream-job/config"
	"github.com/Dima-F/dream-job/internal/home"
	"github.com/Dima-F/dream-job/pkg/logger"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)

	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New())

	home.NewHandler(app, customLogger)
	app.Listen(":3000")
}
