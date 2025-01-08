package main

import (
	"github.com/gofiber/fiber/v2"
)

type Client struct {
	name     string
	email    string
	password string
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Привет, мир!")
	})
	app.Post("/registration", registrationHandler)

	app.Listen(":3000")
}

func (c *Client) registrationHandler() {

}
