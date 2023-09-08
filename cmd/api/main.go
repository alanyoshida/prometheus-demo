package main

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    prometheus := fiberprometheus.New("my-service-name")
    prometheus.RegisterAt(app, "/metrics")
    app.Use(prometheus.Middleware)

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })
    app.Get("/400", func(c *fiber.Ctx) error {
        return c.Status(fiber.StatusBadRequest).SendString("Hello, BadRequest!")
    })
    app.Get("/500", func(c *fiber.Ctx) error {
        return c.Status(fiber.StatusInternalServerError).SendString("Hello, InternalServerError!")
    })
    app.Get("/502", func(c *fiber.Ctx) error {
        return c.Status(fiber.StatusBadGateway).SendString("Hello, BadGateway!")
    })
    app.Listen(":3000")
}