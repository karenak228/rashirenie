package handlers //ручки

import (
	"rashirenie/register-service/internal/requests"

	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
}

func New() *Handlers {
	return &Handlers{}
}

func (h *Handlers) RegistrationHandler(c *fiber.Ctx) error {
	r := &requests.Client{} //
	if err := c.BodyParser(r); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "bed request"})
	}
	return c.Status(200).JSON(fiber.Map{
		"status": "good request"})
}
