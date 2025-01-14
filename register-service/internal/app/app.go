package app

import (
	"rashirenie/register-service/internal/handlers"
	"rashirenie/register-service/internal/repository"
	"rashirenie/register-service/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type App struct {
	services   *services.Services
	handlers   *handlers.Handlers
	app        *fiber.App
	repository *repository.Repositori
}

func New() *App {
	a := &App{}
	db := repository.CreateConnection() //подключенная бд
	a.repository = repository.New(db)
	a.app = fiber.New()
	a.services = services.New(a.repository)
	a.handlers = handlers.New(a.services) //вызвали функцию new из эндпоинт(хендлера)
	a.routers()
	return a
}

func (a *App) routers() {
	a.app.Use(logger.New(), recover.New())
	a.app.Post("/registration", a.handlers.RegistrationHandler)
	a.app.Post("/login", a.handlers.LoginHandler)
}

func (a *App) Run() {
	a.app.Listen(":3000")
}
