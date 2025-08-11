package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	return c.SendString("String answer")
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.customLogger.Info().Bool("IsAdmin", true).Str("Email", "dd@d.d").Int("Age", 35).Msg("Output example")
	return fiber.NewError(fiber.StatusBadRequest, "Limit params is undefined")
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
	}
	api := h.router.Group("/api")
	api.Get("/", h.home)
	api.Get("/error", h.error)
}
