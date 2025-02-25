package middlewares

import (
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"os"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

// CustomClaims, JWT içindeki özel iddiaları tutar.
type CustomClaimsUser struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
	Level    int    `json:"level"` 
}

// GenerateJWT, kullanıcı adına, role ve yetki seviyesine göre JWT oluşturur.
func UserGenerateJWT(username string, role string, level int) (string, error) {
	claims := CustomClaimsUser{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // 2 saat geçerli JWT
			Subject:   username,
			Issuer:    role, // Kullanıcı rolünü Issuer alanına ekleyin
		},
		Username: username,
		Role:     role,
		Level:    level,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// UserParseJWT, verilen token'ı çözümleyerek claims'i döndürür.
func UserParseJWT(tokenString string) (*CustomClaimsUser, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsUser{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaimsUser)
	if !ok {
		return nil, fmt.Errorf("JWT doğrulama hatası")
	}

	return claims, nil
}

// UserMiddleTokenControlLevel, belirli bir yetki seviyesine sahip kullanıcıları kontrol eden middleware.
func UserMiddleTokenControlLevel(requiredLevel int) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// "Baraer" adlı çerezden token alınır.
		cookie := c.Cookies("Baraer")

		if cookie == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"Status": 403,
				"Success": false,
				"Message": "Yetkilendirme Başarısız",
			})
		}

		// JWT'yi çözümleme işlemi yapılır.
		claims, err := UserParseJWT(cookie)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"Status": 400,
				"Success": false,
				"Message": "Geçersiz Token.",
			})
		}

		// Yetki seviyesini kontrol et
		if claims.Level < requiredLevel {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"Status": 401,
				"Success": false,
				"Message": "Yetkisiz erişim.",
			})
		}

		// Middleware'den sonraki işlemleri devam ettir.
		return c.Next()
	}
}

// UserMiddleTokenControl, JWT'yi kontrol eden Fiber middleware.
func UserMiddleTokenControl(c *fiber.Ctx) error {
	// "Baraer" adlı çerezden token alınır.
	cookie := c.Cookies("Baraer")

	if cookie == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Success": false,
			"Message": "Yetkilendirme Başarısız",
		})
	}

	// JWT'yi çözümleme işlemi yapılır.
	_, err := UserParseJWT(cookie)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Success": false,
			"Message": "Geçersiz Token.",
		})
	}

	// Middleware'den sonraki işlemleri devam ettir.
	return c.Next()
}
