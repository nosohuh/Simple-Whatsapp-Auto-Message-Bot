package usercontroller

import (

	"github.com/gofiber/fiber/v2"
	db "github.com/while/payproje/db"
	usermodels "github.com/while/payproje/models/UserModels"
)


func GetFullUser(c *fiber.Ctx) error {
	var users []usermodels.Users
	result := db.DB.Order("created_at desc").Find(&users)

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
		"message": "User data retrieved successfully",
		"data":    users,
	})
}