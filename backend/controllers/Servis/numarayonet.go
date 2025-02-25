package servis

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/while/payproje/db"
	servismodel "github.com/while/payproje/models/ServisModels"
)

func AddNumer(c *fiber.Ctx) error {
	servis := servismodel.Servis{}

	if err := c.BodyParser(&servis); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Boş gönderemezsiniz",
			"error":   map[string]interface{}{},
		})
	} 
	if len(servis.Numara) < 4 {
		return c.Status(400).JSON(fiber.Map{
			"Success": false,
			"Message": "Çok Kısa Değer Girdiniz.",
		})
	}

	LobbyCreate := servismodel.Servis{
		Numara: servis.Numara,
		Mesaj: servis.Mesaj,
	}

	if err := db.DB.Create(&LobbyCreate).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Veritabanına kaydetme hatası")
	}

	return c.Status(200).JSON(fiber.Map{
		"Success": true,
		"Message": "Numara Eklendi.",
		
	})
}


func NumaraGetir(c *fiber.Ctx) error {
	var servis []servismodel.Servis
	result := db.DB.Order("created_at desc").Find(&servis)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"type":    "Error",
			"message": "User data couldn't be retrieved",
		})}
		

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"type":    "Success",
		"data":    servis,
	})
}