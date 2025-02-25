package usercontroller

import (
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	db "github.com/while/payproje/db"
	usermodels "github.com/while/payproje/models/UserModels"
	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode" // QR kod oluşturma kütüphanesi
)

// Kullanıcıya OTP ekleme
func CreateUserOTP(c *fiber.Ctx) error {
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

	var findusername usermodels.Users
	fresult := db.DB.Where("username = ?", usermodel.Username).First(&findusername)

	if fresult.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"Success": false,
			"Message": "Böyle bir kullanıcı yok.",
		})
	}

	// TOTP anahtarını oluştur
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "RoketHavale", // Uygulamanızın adı
		AccountName: usermodel.Username,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("TOTP anahtarı oluşturma hatası")
	}

	// Kullanıcının TOTP anahtarını güncelle
	findusername.TOTPKey = key.Secret() // Anahtarı düz metin olarak kaydet

	if err := db.DB.Save(&findusername).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Veritabanına kaydetme hatası")
	}

	// QR kodu oluştur
	qrCode, err := qrcode.Encode(key.URL(), qrcode.Medium, 256)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("QR kodu oluşturma hatası")
	}

	// QR kodunu base64 formatına çevir
	base64QR := base64.StdEncoding.EncodeToString(qrCode)

	// Başarılı işlem durumu
	return c.Status(200).JSON(fiber.Map{
		"Success": true,
		"Message": "OTP oluşturuldu",
		"TOTPKey": key.Secret(), // Anahtarı kullanıcıya ilet
		"TOTPURL": key.URL(),    // Kullanıcının QR kodu taraması için URL
		"QRCode":  base64QR,     // Base64 formatındaki QR kod
	})
}
