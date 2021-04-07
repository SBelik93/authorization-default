package middlewares

import (
	"114_fiber2_gorm2_nolayers/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Rights(next fiber.Handler, successPermissions ...string) fiber.Handler {
	return func(f *fiber.Ctx) (err error) {
		permissionsH := f.Get("permissions", "")
		if permissionsH == "" {
			models.Error(fmt.Errorf("no permissions"), f, fiber.StatusUnauthorized)
			return
		}
		//var userPermissions []models.JwtPermission
		//err := json.Unmarshal([]byte(permissionsH), &userPermissions)
		//if !models.Error(err, f, models.CodeClientUnauthorized) {
		//	return
		//}
		//
		//for _, successPermission := range successPermissions {
		//	for _, userPermission := range userPermissions {
		//		if userPermission.TechnicalName == successPermission {
		//			next(f)
		//			return
		//		}
		//	}
		//}
		models.Error(fmt.Errorf("no permissions"), f, fiber.StatusUnauthorized)
		return
	}
}
