package servis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// MessageRequest yapısı, mesaj gönderme isteği için kullanılan verileri temsil eder.
type MessageRequest struct {
	Number  string `json:"number"`
	Message string `json:"message"`
}

// SendMessage, mesaj gönderme işlemini gerçekleştiren fonksiyondur.
func SendMessage(req MessageRequest) (map[string]interface{}, error) {
	// Boş alan kontrolü
	if req.Number == "" || req.Message == "" {
		return nil, fmt.Errorf("numara ve mesaj alanları zorunludur")
	}

	// Node.js API'sine mesaj gönderme isteği
	apiURL := "http://localhost:4020/bot-action"
	
	
	payload := map[string]interface{}{
		"action":  "send-message", 
		"number":  req.Number,     
		"message": req.Message,    
	}

	// JSON formatında payload oluşturma
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("payload oluşturulamadı: %v", err)
	}

	// POST isteği gönderme
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("API'ye istek gönderilemedi: %v", err)
	}
	defer resp.Body.Close()

	// API yanıt kodunu kontrol etme
	log.Println("API Yanıt Kodu:", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API yanıtı hatalı: %s", resp.Status)
	}

	// API yanıtını okuma
	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("API yanıtı okunamadı: %v", err)
	}

	// Başarılı işlem yanıtı
	log.Println("API Yanıtı:", response)
	return response, nil
}
