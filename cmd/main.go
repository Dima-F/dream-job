package main

import (
	"strings"
	"time"

	"github.com/Dima-F/dream-job/config"
	"github.com/Dima-F/dream-job/internal/home"
	"github.com/Dima-F/dream-job/internal/vacancy"
	"github.com/Dima-F/dream-job/pkg/database"
	"github.com/Dima-F/dream-job/pkg/logger"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// config
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	dbConfig := config.NewDatabaseConfig()
	customLogger := logger.NewLogger(logConfig)

	engine := html.New("./html", ".html")
	engine.AddFuncMap(map[string]interface{}{
		"ToUpper": func(c string) string {
			return strings.ToUpper(c)
		},
	})

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New())
	app.Static("/public", "./public")

	dbpool := database.CreateDbPool(dbConfig, customLogger)
	defer dbpool.Close()

	storage := postgres.New(postgres.Config{
		DB:         dbpool,
		Table:      "session_storage",
		Reset:      false,
		GCInterval: 10 * time.Second,
	})

	store := session.New(session.Config{
		Storage: storage,
	})

	// repositories
	vacancyRepo := vacancy.NewVacancyRepository(dbpool, customLogger)

	// handlers
	home.NewHandler(app, customLogger, vacancyRepo, store)
	vacancy.NewHandler(app, customLogger, vacancyRepo)

	app.Listen(":3000")
}
