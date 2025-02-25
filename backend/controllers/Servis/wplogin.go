package servis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"log"
)

// WhatsApp QR kodunu almak için handler fonksiyonu
func GetQRCodeHandler(c *fiber.Ctx) error {
	// Node.js API'ye QR kodu almak için istek gönderme
	apiURL := "http://localhost:4020/bot-action" // API'nin doğru adresi

	// Payload'ı oluştur
	payload := map[string]interface{}{
		"action": "qr", // QR kodunu almak için action: 'qr'
	}

	// JSON formatında payload'ı encode et
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Payload oluşturulamadı: %v", err),
		})
	}

	// POST isteği gönderme, payload JSON formatında
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("API'ye istek gönderilemedi: %v", err),
		})
	}
	defer resp.Body.Close()

	// Yanıt kodunu kontrol etme
	if resp.StatusCode != http.StatusOK {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("API yanıtı hatalı: %s", resp.Status),
		})
	}

	// API yanıtını okuma (QR kodu verisi burada alınır)
	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("API yanıtı okunamadı: %v", err),
		})
	}

	// QR kodunu döndürme
	qrCode, ok := response["qrCode"].(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "QR kodu yanıtında bir hata oluştu",
		})
	}

	// QR kodunu içeren başarılı işlem yanıtı döndürme
	return c.JSON(fiber.Map{
		"message": "QR kodu başarıyla alındı",
		"qr":      qrCode,
	})
}


// WhatsApp botunu başlatma handler'ı
func LoginHandler(c *fiber.Ctx) error {
	// Node.js API'ye botu başlatma isteği gönderme
	apiURL := "http://localhost:4020/bot-action" // Bot başlatma API'si

	// Payload'ı oluştur
	payload := map[string]interface{}{
		"action": "start", // Botu başlatma isteği
	}

	// JSON formatında payload'ı encode et
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Payload oluşturulamadı: %v", err),
		})
	}

	// POST isteği gönderme, payload JSON formatında
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("API'ye istek gönderilemedi: %v", err),
		})
	}
	defer resp.Body.Close()

	// Yanıt kodunu kontrol etme
	if resp.StatusCode != http.StatusOK {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("API yanıtı hatalı: %s", resp.Status),
		})
	}

	// API yanıtını okuma
	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("API yanıtı okunamadı: %v", err),
		})
	}

	// API'den gelen yanıtla birlikte başarılı işlem yanıtı döndürme
	return c.JSON(fiber.Map{
		"message":  "Node.js API'ye istek başarılı, bot başlatıldı!",
		"response": response,
	})
}

func StopHandler(c *fiber.Ctx) error {
	// Node.js API'ye botu başlatma isteği gönderme
	apiURL := "http://localhost:4020/bot-action" // Bot başlatma API'si

	// Payload'ı oluştur
	payload := map[string]interface{}{
		"action": "stop", // Botu başlatma isteği
	}

	// JSON formatında payload'ı encode et
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Payload oluşturulamadı: %v", err),
		})
	}

	// POST isteği gönderme, payload JSON formatında
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("API'ye istek gönderilemedi: %v", err),
		})
	}
	defer resp.Body.Close()

	// Yanıt kodunu kontrol etme
	if resp.StatusCode != http.StatusOK {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("API yanıtı hatalı: %s", resp.Status),
		})
	}

	// API yanıtını okuma
	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("API yanıtı okunamadı: %v", err),
		})
	}

	// API'den gelen yanıtla birlikte başarılı işlem yanıtı döndürme
	return c.JSON(fiber.Map{
		"message":  "Node.js API'ye istek başarılı, bot başlatıldı!",
		"response": response,
	})
}

func SendMessageHandler(c *fiber.Ctx) error {
	// Kullanıcıdan gelen verileri okuma
	type MessageRequest struct {
		Number  string `json:"number"`
		Message string `json:"message"`
	}
	var req MessageRequest
	if err := c.BodyParser(&req); err != nil {
		// Geçersiz veri formatı hatası
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Geçersiz veri formatı: %v", err),
		})
	}

	// Boş alan kontrolü
	if req.Number == "" || req.Message == "" {
		// Numara ve mesaj zorunlu alanlar
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Numara ve mesaj alanları zorunludur!",
		})
	}

	// Node.js API'sine mesaj gönderme isteği
	apiURL := "http://localhost:4020/bot-action" // Mesaj gönderme API'si

	// Payload oluştur
	payload := map[string]interface{}{
		"action":  "send-message", // Mesaj gönderme işlemi
		"number":  req.Number,     // Hedef numara
		"message": req.Message,    // Mesaj içeriği
	}

	// JSON formatında payload oluşturma
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		// Payload oluşturulamadı hatası
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Payload oluşturulamadı: %v", err),
		})
	}

	// POST isteği gönderme
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		// API'ye istek gönderilemedi hatası
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("API'ye istek gönderilemedi: %v", err),
		})
	}
	defer resp.Body.Close()

	// API yanıt kodunu kontrol etme
	log.Println("API Yanıt Kodu:", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("API yanıtı hatalı: %s", resp.Status),
		})
	}

	// API yanıtını okuma
	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		// API yanıtı okunamadı hatası
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("API yanıtı okunamadı: %v", err),
		})
	}

	// API yanıtını loglama
	log.Println("API Yanıtı:", response)

	// Başarılı işlem yanıtı
	return c.JSON(fiber.Map{
		"message":  "Mesaj başarıyla gönderildi!",
		"response": response,
	})
}

func BotStatusHandler(c *fiber.Ctx) error {
	// Hedef API URL
	apiURL := "http://localhost:4020/bot-status" 

	// GET isteği gönderme
	resp, err := http.Get(apiURL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("API'ye istek gönderilemedi: %v", err),
		})
	}
	defer resp.Body.Close()

	// Yanıtı okuma
	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("API yanıtı okunamadı: %v", err),
		})
	}

	// API'den gelen yanıtla birlikte başarılı işlem yanıtı döndürme
	return c.JSON(fiber.Map{
		"message":  "Bot durumu başarıyla alındı!",
		"response": response,
	})
}
