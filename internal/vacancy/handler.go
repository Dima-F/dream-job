package vacancy

import (
	"time"

	"github.com/Dima-F/dream-job/pkg/tadapter"
	"github.com/Dima-F/dream-job/pkg/validator"
	"github.com/Dima-F/dream-job/views/components"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
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
	form := VacancyCreateForm{
		Email: c.FormValue("email"),
	}
	errors := validate.Validate(
		&validators.EmailIsPresent{
			Name:    "Email",
			Field:   form.Email,
			Message: "Email is empty or not valid",
		},
	)

	time.Sleep(time.Second * 2)

	if len(errors.Errors) > 0 {
		failComponent := components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tadapter.Render(c, failComponent)
	}

	successComponent := components.Notification("Vacancy created successfully", components.NotificationSuccess)
	return tadapter.Render(c, successComponent)
}
