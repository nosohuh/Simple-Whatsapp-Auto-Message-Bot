package usercontroller

import (
	"github.com/gofiber/fiber/v2"
	midllewares "github.com/while/payproje/midllewares"
	"time"
)

func UserLogout(c *fiber.Ctx) error {
	cookie := c.Cookies("Baraer")
	if cookie == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Success": false,
			"Message": "Oturum bulunamadı",
		})
	}
	token, err := midllewares.UserGenerateJWT("exp", "exp", 3600) // Assuming 3600 as the third argument
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "JWT convert error",
		})
	}

	// JWT'yi silme
	c.ClearCookie("Baraer")

	// Yeni çerezi ayarla
	c.Cookie(&fiber.Cookie{
		Name:    "Baraer",
		Value:   token,
		Expires: time.Now().Add(-time.Hour), // Geçmişe ayarlayarak siler
		Path:     "/",
		HTTPOnly: true,   // HTTPOnly ayarı, JavaScript tarafından erişimine izin vermek istiyorsanız false olarak ayarlanabilir
		Secure:   false,  // Güvenli bağlantı (HTTPS) gerekiyorsa true olarak ayarlanabilir
		SameSite: "none", // SameSite ayarı, gereksinimlerinize göre ayarlanabilir
	})

	// Kullanıcıyı çıkış sayfasına yönlendirme
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Success": false,
		"Message": "İstek başarılı, çıkış yapılıyor.",
	})
}
