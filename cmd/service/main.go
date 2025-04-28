package main

import "github.com/gofiber/fiber/v2"
import "github.com/grzadr/refscaler/units"



func main() {
	app := fiber.New()

	app.Get("/units", func(c *fiber.Ctx) error {
		return c.SendString(units.EmbeddedUnitRegistry.ToJSON())
	})

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
