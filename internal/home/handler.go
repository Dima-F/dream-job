package home

import (
	"net/http"

	"github.com/Dima-F/dream-job/internal/vacancy"
	"github.com/Dima-F/dream-job/pkg/tadapter"
	"github.com/Dima-F/dream-job/views"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
	repository   *vacancy.VacancyRepository
}

type User struct {
	Id   int
	Name string
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	vacancies, err := h.repository.GetAll()
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		return c.SendStatus(500)
	}
	component := views.Main(vacancies)
	return tadapter.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.customLogger.Info().Bool("IsAdmin", true).Str("Email", "dd@d.d").Int("Age", 35).Msg("Output example")
	return fiber.NewError(fiber.StatusBadRequest, "Limit params is undefined")
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger, repository *vacancy.VacancyRepository) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
		repository:   repository,
	}
	api := h.router.Group("/api")
	api.Get("/", h.home)
	api.Get("/error", h.error)
}
