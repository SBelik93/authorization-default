package v1

import (
	"114_fiber2_gorm2_nolayers/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func SignIn(f *fiber.Ctx) (err error){
	var (
		cred = models.BodyCredentials{}
		user models.User
	)

	err = f.BodyParser(&cred)
	if !models.ParsingError(err, f,  fiber.StatusBadRequest) {
		return
	}

	err = user.LoginCheck(cred)
	if !models.Error(err, f, fiber.StatusUnauthorized) {
		return
	}

	token, err := user.CreateJwtToken()
	if !models.Error(err, f,  fiber.StatusUnauthorized) {
		return
	}

	f.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	user.ClaimToken = token
	res := models.Response{
		Status: true,
		Result: user,
	}
	err = f.JSON(res)
	return
}


