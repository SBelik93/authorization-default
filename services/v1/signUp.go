package v1

import (
	"114_fiber2_gorm2_nolayers/models"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(f *fiber.Ctx) (err error) {
	var (
		user models.User
		tx = models.GetDB().WithContext(context.Background()).Begin()
	)

	err = f.BodyParser(&user)
	if !models.ParsingError(err, f, fiber.StatusBadRequest) {
		tx.Rollback()
		return
	}

	err = user.Validate()
	if !models.Error(err, f, fiber.StatusBadRequest) {
		tx.Rollback()
		return
	}
	hashedPassword, err  := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if !models.Error(err, f, fiber.StatusNotAcceptable) {
		tx.Rollback()
		return
	}
	user.Password = string(hashedPassword)

	err = tx.Create(&user).Error
	if !models.Error(err, f, fiber.StatusNotAcceptable) {
		tx.Rollback()
		return
	}

	token, err := user.CreateJwtToken()
	if !models.Error(err, f, fiber.StatusNotAcceptable) {
		tx.Rollback()
		return
	}

	f.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	user.ClaimToken = token
	res := models.Response{
		Status:  true,
		Result:  user,
	}
	err = f.JSON(res)
	if !models.Error(err, f, fiber.StatusNotAcceptable) {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}
