package services

import (
	v1 "114_fiber2_gorm2_nolayers/services/v1"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Static("/swagger", "swagger")
	app.Get("/health_check", HealthCheck)

	v1.Use(app)
}
