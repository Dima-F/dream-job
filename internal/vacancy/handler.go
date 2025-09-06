package vacancy

import (
	"github.com/Dima-F/dream-job/pkg/tadapter"
	"github.com/Dima-F/dream-job/views/components"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type VacancyHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &VacancyHandler{
		router:       router,
		customLogger: customLogger,
	}
	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.createVacancy)
}

func (h *VacancyHandler) createVacancy(c *fiber.Ctx) error {
	email := c.FormValue("email")
	// h.customLogger.Info().Msg(email)

	if email == "" {
		failComponent := components.Notification("Something get wrong", components.NotificationFail)
		return tadapter.Render(c, failComponent)
	}

	successComponent := components.Notification("Vacancy created successfully", components.NotificationSuccess)
	return tadapter.Render(c, successComponent)
}
