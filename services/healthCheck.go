package services

import (
	"114_fiber2_gorm2_nolayers/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) (err error) {
	fmt.Println("--------------health-check----------------")
	var response = models.Response{
		Status:  true,
		Message: "success",
		Result: map[string]interface{}{
			"1": "I'm",
			"2": "ready",
			"3": true,
		},
	}
	err = c.JSON(response)
	if err != nil {
		panic(err)
	}
	return
}
