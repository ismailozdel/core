package middlewares

import (
	"github.com/atolye/cursorgotemplate/core/database"
	"github.com/gofiber/fiber/v2"
)

func databaseSelect(c *fiber.Ctx) error {

	c.Locals("db", database.DB)

	// usr := c.Locals("user").(*models.Claims)

	// db, err := database.GetCompanyDB(usr.CompanyID)
	// if err != nil {
	// 	return httputils.PrepareInternalServerError(err.Error())
	// }
	// if db == nil {
	// 	return httputils.PrepareInternalServerError("Database not found")
	// }

	// c.Locals("db", db)

	return c.Next()
}
