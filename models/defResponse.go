package models

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

//swagger:model http_response
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func Error(err error, c *fiber.Ctx, code int) bool {
	if err != nil {
		res := Response{
			Status:  false,
			Message: err.Error(),
		}
		_ = c.SendStatus(code)
		err = c.JSON(res)
		if err != nil {
			fmt.Println(Red, "PAAANIIIIICA-AAAAAaaa: ", err.Error(), Reset)
		}
		return false
	}
	return true
}

func ParsingError(err error, c *fiber.Ctx, code int) (status bool) {
	if err != nil {
		res := Response{
			Status:  false,
			Message: fmt.Sprintf("%s, %v", ErrorParsingBody, err.Error()),
		}
		c.SendStatus(code)
		err = c.JSON(res)
		if err != nil {
			fmt.Println(Red, "PAAANIIIIICA-AAAAAaaa: ", err.Error(), Reset)
		}
		return false
	}
	return true
}
