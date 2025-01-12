package handlers //ручки

import (
	"rashirenie/register-service/internal/requests"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var validate = validator.New()

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
			"status": "bad request"})
	}

	errors := validate.Struct(r)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(errors.Error())
	}

	logrus.Info("Все ок")

	return c.Status(200).JSON(fiber.Map{
		"status": "good request"})
}

func (h *Handlers) LoginHandler(c *fiber.Ctx) error {
	r := requests.Login{}
	if err := c.BodyParser(&r); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "bad request"})
	}

	errors := validate.Struct(r)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(errors.Error())
	}

	err := h.servises.Login(r)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "bad request"})
	}

	logrus.Info("Все ок")

	return c.Status(200).JSON(fiber.Map{
		"status": "good request"})
}
