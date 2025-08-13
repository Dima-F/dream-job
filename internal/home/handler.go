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
	// teml := template.Must(template.ParseFiles("./html/page.html"))
	// data := struct{ Count int }{Count: 1}
	// var tpl bytes.Buffer
	// if err := teml.Execute(&tpl, data); err != nil {
	// 	return fiber.NewError(fiber.StatusBadRequest, "Template compile error")
	// }
	// c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	// return c.Send(tpl.Bytes())
	return c.Render("page", fiber.Map{
		"Count":   3,
		"IsAdmin": true,
		"CanUse":  false,
	})
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
