package home

import (
	"math"
	"net/http"

	"github.com/Dima-F/dream-job/internal/vacancy"
	"github.com/Dima-F/dream-job/pkg/tadapter"
	"github.com/Dima-F/dream-job/views"
	"github.com/Dima-F/dream-job/views/components"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
	repository   *vacancy.VacancyRepository
	store        *session.Store
}

type User struct {
	Id   int
	Name string
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	PAGE_ITEMS := 5
	page := c.QueryInt("page", 1)

	sess, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}

	if name, ok := sess.Get("name").(string); ok {
		h.customLogger.Info().Msg(name)
	}
	count := h.repository.CountAll()
	vacancies, err := h.repository.GetAll(PAGE_ITEMS, (page-1)*PAGE_ITEMS)
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		return c.SendStatus(500)
	}
	component := views.Main(vacancies, int(math.Ceil(float64(count/PAGE_ITEMS))), page)
	return tadapter.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.customLogger.Info().Bool("IsAdmin", true).Str("Email", "dd@d.d").Int("Age", 35).Msg("Output example")
	return fiber.NewError(fiber.StatusBadRequest, "Limit params is undefined")
}

func (h *HomeHandler) login(c *fiber.Ctx) error {
	component := views.Login()
	return tadapter.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) apiLogin(c *fiber.Ctx) error {
	form := LoginForm{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	if form.Email == "a@a.ua" && form.Password == "123456" {
		sess, err := h.store.Get(c)
		if err != nil {
			panic(err)
		}
		sess.Set("email", form.Email)
		if err := sess.Save(); err != nil {
			panic(err)
		}
		return c.Redirect("/", http.StatusOK)
	}
	errComponent := components.Notification("Login error", components.NotificationFail)
	return tadapter.Render(c, errComponent, http.StatusBadRequest)
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger, repository *vacancy.VacancyRepository, store *session.Store) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
		repository:   repository,
		store:        store,
	}
	api := h.router.Group("/api")
	api.Get("/", h.home)
	api.Get("/login", h.login)
	api.Post("/login", h.apiLogin)
	api.Get("/error", h.error)
}
