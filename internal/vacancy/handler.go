package vacancy

import (
	"net/http"

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
	repository   *VacancyRepository
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger, repository *VacancyRepository) {
	h := &VacancyHandler{
		router:       router,
		customLogger: customLogger,
		repository:   repository,
	}
	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.createVacancy)
}

func (h *VacancyHandler) createVacancy(c *fiber.Ctx) error {
	form := VacancyCreateForm{
		Email:    c.FormValue("email"),
		Location: c.FormValue("location"),
		Type:     c.FormValue("type"),
		Company:  c.FormValue("company"),
		Role:     c.FormValue("role"),
		Salary:   c.FormValue("salary"),
	}
	errors := validate.Validate(
		&validators.EmailIsPresent{
			Name:    "Email",
			Field:   form.Email,
			Message: "Email is empty or not valid",
		},
		&validators.StringIsPresent{
			Name:    "Location",
			Field:   form.Location,
			Message: "Location is required",
		},
		&validators.StringIsPresent{
			Name:    "Type",
			Field:   form.Type,
			Message: "Company type is required",
		},
		&validators.StringIsPresent{
			Name:    "Company",
			Field:   form.Company,
			Message: "Company name is required",
		},
		&validators.StringIsPresent{
			Name:    "Role",
			Field:   form.Role,
			Message: "Role is required",
		},
		&validators.StringIsPresent{
			Name:    "Salary",
			Field:   form.Salary,
			Message: "Salary is required",
		},
	)

	// time.Sleep(time.Second * 2)

	if len(errors.Errors) > 0 {
		failComponent := components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tadapter.Render(c, failComponent, http.StatusBadRequest)
	}

	err := h.repository.addVacancy(form)

	if err != nil {
		errComponent := components.Notification(err.Error(), components.NotificationFail)
		return tadapter.Render(c, errComponent, http.StatusBadRequest)
	}

	successComponent := components.Notification("Vacancy created successfully", components.NotificationSuccess)
	return tadapter.Render(c, successComponent, http.StatusOK)
}
