package v1

import (
	"github.com/gofiber/fiber/v2"
)

func Use(app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Post("/api/v1/sign-in", SignIn)
	v1.Post("/api/v1/sign-up", SignUp)
}
