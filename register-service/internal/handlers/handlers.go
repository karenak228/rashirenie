package handlers //ручки

import "github.com/gofiber/fiber/v2"

type Handlers struct {
}

func New() *Handlers {
	return &Handlers{}
}

func (h *Handlers) RegistrationHandler(c *fiber.Ctx) error {
	return nil
}
