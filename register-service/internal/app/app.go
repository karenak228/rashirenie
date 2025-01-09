package app

import (
	"rashirenie/register-service/internal/handlers"
	"rashirenie/register-service/internal/services"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	services *services.Services
	handlers *handlers.Handlers
	app      *fiber.App
}

func New() *App {
	a := &App{}
	a.app = fiber.New()
	a.services = services.New()
	a.handlers = handlers.New()
	a.routers()
	return a
}

func (a *App) routers() {
	a.app.Post("/registration", a.handlers.RegistrationHandler)
}

func (a *App) Run() {
	a.app.Listen(":3000")
}
