package home

import "github.com/gofiber/fiber/v2"

type HomeHandler struct {
	router fiber.Router
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	return c.SendString("String answer")
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	// return fiber.ErrBadRequest
	return fiber.NewError(fiber.StatusBadRequest, "Limit params is undefined")
}

func NewHandler(router fiber.Router) {
	h := &HomeHandler{router: router}
	api := h.router.Group("/api")
	api.Get("/", h.home)
	api.Get("/error", h.error)
}
