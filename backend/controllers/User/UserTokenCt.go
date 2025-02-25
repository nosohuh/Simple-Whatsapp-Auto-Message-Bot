package usercontroller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	midllewares "github.com/while/payproje/midllewares"
)

func TokenConrol(c *fiber.Ctx) error {
	cookie := c.Cookies("Baraer")
	if cookie == "" {
		return c.Status(400).JSON(fiber.Map{
			"Success": false,
			"Message": "Yetkilendirme Başarısız",
		})
	}

	claims, err := midllewares.UserParseJWT(cookie)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Success": false,
			"Message": "Geçersiz Token.",
		})
	}
	return c.JSON(fiber.Map{
		"status":   200,
		"operator": claims.Subject,
		"username":     claims.Username,
		"role":     claims.Role, // Rol bilgisini claims.Role alanından alıyoruz.
		"level":    claims.Level,
		"message":  fmt.Sprintf("Hoş geldiniz, %s!", claims.Subject),
	})
}



