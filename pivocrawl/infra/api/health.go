package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func AddHealthRoutes(app *fiber.App) {
	app.Get("/health", getHealth)
}

func getHealth(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
