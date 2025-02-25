package usercontroller

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/while/payproje/db"
	midllewares "github.com/while/payproje/midllewares"
	usermodels "github.com/while/payproje/models/UserModels"
	"golang.org/x/crypto/bcrypt"
	"github.com/pquerna/otp/totp"
	"time"
)


func Login(c *fiber.Ctx) error {
	usermodel := usermodels.Users{}

	if err := c.BodyParser(&usermodel); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Username or password are not valid",
			"error":   nil,
		})
	}

	if len(usermodel.Username) < 4 || len(usermodel.Password) < 4 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Username or password are too short",
		})
	}

	var findusername usermodels.Users
	fresult := db.DB.Where("username = ?", usermodel.Username).First(&findusername)

	if fresult.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Username is not valid",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(findusername.Password), []byte(usermodel.Password))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Password is incorrect",
		})
	}

	// TOTP anahtarı kontrol ediliyor
	if findusername.TOTPKey == "" {
		// OTP boşsa, normal giriş işlemleri
		token, err := midllewares.UserGenerateJWT(findusername.Username, findusername.Role, findusername.Level)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "JWT oluşturma hatası",
			})
		}

		expire := time.Now().AddDate(0, 0, 1)
		cookie := fiber.Cookie{
			Name:     "Baraer",
			Value:    token,
			Expires:  expire,
			Path:     "/",
			HTTPOnly: true,   // HTTPOnly ayarı
			Secure:   false,  // Güvenli bağlantı (HTTPS) gerekiyorsa true olarak ayarlanabilir
			SameSite: "none", // SameSite ayarı
		}

		c.Cookie(&cookie)

		return c.JSON(fiber.Map{
			"success":  true,
			"status":   200,
			"message":  "Giriş başarılı.",
			"username": findusername.Username,
		})
	}

	// OTP doluysa, OTP doğrulaması isteniyor
	return c.JSON(fiber.Map{
		"success":  false,
		"status":   400,
		"message":  "Lütfen OTP doğrulamasını yapın.",
		"username": usermodel.Username,
	})
}

// TOTP kodunu doğrulama işlemi
func LoginVerify(c *fiber.Ctx) error {
	type VerifyRequest struct {
		Username string `json:"username"`
		Code     string `json:"code"`
	}

	var req VerifyRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Username and code are required",
		})
	}

	var user usermodels.Users
	if err := db.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Böyle bir kullanıcı yok.",
		})
	}

	valid := totp.Validate(req.Code, user.TOTPKey)
	if !valid {
		return c.Status(403).JSON(fiber.Map{
			"success": false,
			"message": "OTP kodu geçersiz",
		})
	}

	token, err := midllewares.UserGenerateJWT(user.Username, user.Role, user.Level)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "JWT oluşturma hatası",
		})
	}


	expire := time.Now().AddDate(0, 0, 1)
	cookie := fiber.Cookie{
		Name:     "Baraer",
		Value:    token,
		Expires:  expire,
		Path:     "/",
		HTTPOnly: true,   
		Secure:   false, 
		SameSite: "none", 
	}

	c.Cookie(&cookie)

	// JWT oluşturma başarılıysa
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "2FA doğrulaması başarılı.",
		"token":   token,
	})
}