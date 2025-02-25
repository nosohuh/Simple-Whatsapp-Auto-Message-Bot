package usercontroller

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/while/payproje/db"
	midllewares "github.com/while/payproje/midllewares"
	usermodels "github.com/while/payproje/models/UserModels"
	"golang.org/x/crypto/bcrypt"
)

// JWT ile giriş yapma ve yeni kullanıcı oluşturma
func CreateUser(c *fiber.Ctx) error {
	usermodel := usermodels.Users{}

	if err := c.BodyParser(&usermodel); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Boş gönderemezsiniz",
			"error":   map[string]interface{}{},
		})
	}

	if len(usermodel.Username) < 4 {
		return c.Status(400).JSON(fiber.Map{
			"Success": false,
			"Message": "Kullanıcı adı çok kısa",
		})
	}

	if len(usermodel.Password) < 4 {
		return c.Status(400).JSON(fiber.Map{
			"Success": false,
			"Message": "Şifre çok kısa",
		})
	}

	var findusername usermodels.Users
	fresult := db.DB.Where("username = ?", usermodel.Username).First(&findusername)

	if fresult.Error == nil {
		return c.Status(400).JSON(fiber.Map{
			"Success": false,
			"Message": "Kullanıcı adı zaten var",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usermodel.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Şifre hash oluşturma hatası")
	}

	token, err := midllewares.UserGenerateJWT(usermodel.Username, usermodel.Role, usermodel.Level)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("JWT oluşturma hatası")
	}


	LobbyCreate := usermodels.Users{
		Username: usermodel.Username,
		Password: string(hashedPassword),
		Level:    usermodel.Level,
		Role:     usermodel.Role,
	}

	if err := db.DB.Create(&LobbyCreate).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Veritabanına kaydetme hatası")
	}


	// Başarılı işlem durumu
	return c.Status(200).JSON(fiber.Map{
		"Success": true,
		"Message": "Kullanıcı oluşturuldu ve JWT oluşturuldu.",
		"Token":   token,
		
	})
}
