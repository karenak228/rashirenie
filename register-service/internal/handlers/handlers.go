package handlers //ручки

import (
	"rashirenie/register-service/internal/requests"

	"github.com/gofiber/fiber/v2"
)

type Services interface {
	Login(requests.Login) error
}

type Handlers struct {
	servises Services
}

func New(s Services) *Handlers {

	return &Handlers{servises: s}
}

func (h *Handlers) RegistrationHandler(c *fiber.Ctx) error {
	r := &requests.Client{}
	if err := c.BodyParser(r); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "bed request"})
	}
	return c.Status(200).JSON(fiber.Map{
		"status": "good request"})
}

func (h *Handlers) LoginHandler(c *fiber.Ctx) error {
	r := requests.Login{}
	if err := c.BodyParser(&r); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "bed request"})
	}
	err := h.servises.Login(r)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "bed request"})
	}
	return c.Status(200).JSON(fiber.Map{
		"status": "good request"})
}
