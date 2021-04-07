package main

import (
	"114_fiber2_gorm2_nolayers/models"
	"114_fiber2_gorm2_nolayers/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	env, err := models.NewEnvironment()
	if err != nil {
		panic(err)
	}
	sqlDB, err := models.NewDB(&env)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = sqlDB.Close()
	}()

	recoverCfg := fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			_ = ctx.SendStatus(fiber.StatusInternalServerError)
			_ = ctx.JSON(fiber.Map{
				"status":  false,
				"message": err.Error(),
			})
			return nil
		},
	}

	app := fiber.New(recoverCfg)
	if env.Contour == "local" {
		fmt.Printf("health check to localhost%s%s \n", env.Port, "/health_check")
		fmt.Printf("documentations to localhost%s%s \n", env.Port, "/swagger/")
	} else {
		fmt.Printf("health check to started-server-ip%s%s \n", env.Port, "/health_check")
		fmt.Printf("documentations to  started-server-ip%s%s \n", env.Port, "/swagger/")
	}
	app.Use(cors.New(), recover.New())
	services.Routes(app)

	err = app.Listen(env.Port)
	if err != nil {
		panic(err)
	}
}
