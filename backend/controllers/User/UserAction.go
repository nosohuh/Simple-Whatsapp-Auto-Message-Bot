package usercontroller

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/while/payproje/db"
	usermodels "github.com/while/payproje/models/UserModels"
	"golang.org/x/crypto/bcrypt"
)

// UpdateDetailsDealer, bayi (dealer) güncelleme isteği için kullanılan veri modeli.
type UpdateDetailsDealer struct {
	Action     string `json:"action"`
	EditOption string `json:"edit_option"`
	Newdata    string `json:"new_data"`
	UserID     string `json:"id"`
	Role       string `json:"role"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

// DealerFullUpdate, bayi (dealer) bilgilerini güncelleyen API endpoint'i.
func UserFullUpdate(c *fiber.Ctx) error {
	// Bayi (dealer) güncelleme isteğini içeren veri modeli.
	bankup := UpdateDetailsDealer{}
	// İstek gövdesini parse et ve hata kontrolü yap.
	if err := c.BodyParser(&bankup); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"message": "Invalid request body. Please provide valid data.",
			"error":   nil,
		})
	}

	// Eğer action "add" ise, userID kontrolüne gerek yok.
	if bankup.Action != "add" {
		if bankup.Action == "" || bankup.UserID == "" {
			return c.Status(400).JSON(fiber.Map{
				"status":  400,
				"success": false,
				"message": "Action or DealerID is missing.",
				"error":   nil,
			})
		}

		// Bayi (dealer) ID'sine göre arama.
		var finddealer usermodels.Users
		result := db.DB.Where("id = ?", bankup.UserID).First(&finddealer)
		// Hata kontrolü.
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  400,
				"success": true,
				"type":    "Error",
				"message": "Dealer not found.",
				"data":    nil,
			})
		}

		// Bayi (dealer) güncelleme işlemleri.
		switch bankup.Action {
		case "edit":
			return handleEditAction(c, bankup, finddealer)
		case "add":
			return handleCreateAction(c, bankup)
		case "delete":
			return handleDeleteAction(c, finddealer)
		default:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  400,
				"success": true,
				"type":    "Error",
				"message": "Invalid action.",
				"data":    nil,
			})
		}
	} else {
		// "add" eylemi için doğrudan create işlemi gerçekleştirilir.
		return handleCreateAction(c, bankup)
	}

}

func handleEditAction(c *fiber.Ctx, bankup UpdateDetailsDealer, finddealer usermodels.Users) error {
	if bankup.EditOption != "" && bankup.Newdata != "" {
		switch bankup.EditOption {
		case "role":
			return handleEditRole(c, bankup, finddealer)
		case "username":
			return handleEditUsername(c, bankup, finddealer)
		case "password":
			return handleEditPassword(c, bankup, finddealer)
		default:
			return c.Status(400).JSON(fiber.Map{
				"status":  400,
				"success": false,
				"type":    "Error",
				"message": "Invalid EditOption. Please provide a valid field to edit.",
				"data":    nil,
			})
		}
	}
	return c.Status(400).JSON(fiber.Map{
		"status":  400,
		"success": false,
		"type":    "Error",
		"message": "Invalid EditOption. Please provide a valid field to edit.",
		"data":    nil,
	})
}

func handleEditRole(c *fiber.Ctx, bankup UpdateDetailsDealer, finddealer usermodels.Users) error {
	if bankup.Role != "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"type":    "Error",
			"message": "New data conflicts with existing values. Please provide a different value.",
			"data":    nil,
		})
	}

	updateFields := map[string]interface{}{
		"role": bankup.Newdata,
	}

	db.DB.Model(&finddealer).Updates(updateFields)
	if db.DB.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"type":    "Error",
			"message": "Error updating user information.",
			"data":    nil,
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"type":    "Success",
		"message": "Dealer role updated successfully.",
	})
}

func handleEditUsername(c *fiber.Ctx, bankup UpdateDetailsDealer, finddealer usermodels.Users) error {
	updateFields := map[string]interface{}{
		"username": bankup.Newdata,
	}

	db.DB.Model(&finddealer).Updates(updateFields)
	if db.DB.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"type":    "Error",
			"message": "Error updating user information.",
			"data":    nil,
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"type":    "Success",
		"message": "Dealer username updated successfully.",
	})
}

func handleEditPassword(c *fiber.Ctx, bankup UpdateDetailsDealer, finddealer usermodels.Users) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(bankup.Newdata), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error hashing the password")
	}

	updateFields := map[string]interface{}{
		"password": string(hashedPassword),
	}

	db.DB.Model(&finddealer).Updates(updateFields)
	if db.DB.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"type":    "Error",
			"message": "Error updating user information.",
			"data":    nil,
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"type":    "Success",
		"message": "Dealer password updated successfully.",
	})
}

func handleCreateAction(c *fiber.Ctx, bankup UpdateDetailsDealer) error {
	if len(bankup.Username) < 4 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Username is too short.",
		})
	}
	if len(bankup.Password) < 4 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Password is too short.",
		})
	}

	var findusername usermodels.Users
	fresult := db.DB.Where("username = ?", bankup.Username).First(&findusername)

	if fresult.Error == nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Username is already taken.",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(bankup.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error hashing the password")
	}
	LobbyCreate := usermodels.Users{
		Username: bankup.Username,
		Password: string(hashedPassword),
		Role:     bankup.Role,
	}

	if err := db.DB.Create(&LobbyCreate).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error saving to the database")
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "User created.",
	})
}

func handleDeleteAction(c *fiber.Ctx, finddealer usermodels.Users) error {
	db.DB.Delete(&finddealer, "id = ?", finddealer.ID)
	if db.DB.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"success": true,
			"type":    "Error",
			"message": "Error deleting dealer.",
			"data":    nil,
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"type":    "Success",
		"message": "Dealer deleted successfully: " + finddealer.Username,
		"data":    nil,
	})
}
