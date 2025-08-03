package home

import "github.com/gofiber/fiber/v2"

type HomeHandler struct {
	router fiber.Router
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	return c.SendString("String answer")
}

func NewHandler(router fiber.Router) {
	h := &HomeHandler{router: router}
	h.router.Get("/", h.home)
}
