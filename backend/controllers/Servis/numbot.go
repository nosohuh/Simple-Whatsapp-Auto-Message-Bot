package servis

import (
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	db "github.com/while/payproje/db"
	servismodel "github.com/while/payproje/models/ServisModels"
)

var (
	isProcessing bool        // İşlem durumu kontrol değişkeni
	mutex        sync.Mutex  // Mutex ile eşzamanlılık kontrolü
)
func StatusBot(c *fiber.Ctx) error {
	// İşlem durumu kontrolü
	if isProcessing {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Bot çalışıyor.",
			"status":  true,
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Bot şu anda çalışmıyor.",
			"status":  false,
		})
	}
}

func NumBot(c *fiber.Ctx) error {
	mutex.Lock()
	if isProcessing {
		mutex.Unlock()
		log.Println("Bot zaten çalışıyor.")  // Log ekledik
		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
			"success": false,
			"message": "Hâlihazırda bir işlem çalışıyor. Lütfen daha sonra tekrar deneyin.",
		})
	}
	isProcessing = true
	mutex.Unlock()

	log.Println("Bot işlemi başlatıldı.")  // Log ekledik

	var wg sync.WaitGroup  // WaitGroup ile asenkron işlemlerin tamamlanmasını bekle

	defer func() {
		wg.Wait()  // Asenkron işlemlerin tamamlanmasını bekle
		mutex.Lock()
		isProcessing = false  // İşlem tamamlandıktan sonra false yapıyoruz
		mutex.Unlock()
		log.Println("Bot işlemi tamamlandı.")  // Log ekledik
	}()

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ID parametresi eksik",
		})
	}

	// Servis modelini sorgula ve mesajı al
	servis := servismodel.Servis{}
	if err := db.DB.Where("id = ?", id).First(&servis).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Kayıt bulunamadı",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Veritabanı hatası",
			"error":   err.Error(),
		})
	}

	// Veritabanından gelen mesaj
	message := servis.Mesaj // Mesajı veritabanından al

	// Numaraları satırlara böl
	numbers := strings.Split(servis.Numara, "\n")

	// Asenkron işleme başlat
	wg.Add(1)  // Asenkron işlem için WaitGroup sayacını artırıyoruz
	go func() {
		defer wg.Done()  // Asenkron işlem tamamlandığında Done() çağrılır
		log.Println("Asenkron işlem başladı.")  // Log ekledik
		for _, number := range numbers {
			number = strings.TrimSpace(number) // Numara çevresindeki boşlukları temizle
			if number == "" {
				continue // Boş satırları atla
			}

			// "90" ekle
			number = "90" + number

			// Mesaj gönder
			req := MessageRequest{
				Number:  number,
				Message: message,
			}
			if _, err := SendMessage(req); err != nil {
				log.Printf("Mesaj gönderme hatası: %v\n", err)
			}

			// 5-15 saniye arasında rastgele bir süre bekle
			time.Sleep(time.Duration(rand.Intn(11)+5) * time.Second)
		}
		log.Println("Asenkron işlem tamamlandı.")  // Log ekledik
	}()

	// İşlem başlatıldığına dair yanıt döndür
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Numaralar işlenmeye başlandı.",
	})
}
