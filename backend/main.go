package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	db "github.com/while/payproje/db"
	routes "github.com/while/payproje/routes"
	"os"
)

func main() {

	fmt.Println("Program Çalışıyor...")

	var ConfigDefault = helmet.Config{
		XSSProtection:             "0",
		ContentTypeNosniff:        "nosniff",
		XFrameOptions:             "SAMEORIGIN",
		ReferrerPolicy:            "no-referrer",
		CrossOriginEmbedderPolicy: "require-corp",
		CrossOriginOpenerPolicy:   "same-origin",
		CrossOriginResourcePolicy: "same-origin",
		OriginAgentCluster:        "?1",
		XDNSPrefetchControl:       "off",
		XDownloadOptions:          "noopen",
		XPermittedCrossDomain:     "none",
	}
	db.Connect()
	app := fiber.New(fiber.Config{
		Prefork:       false,           // Prefork özelliğini etkinleştir
		CaseSensitive: true,            // URL'lerin büyük/küçük harf duyarlılığını etkinleştir
		StrictRouting: true,            // Kesin yönlendirme modunu etkinleştir
		ServerHeader:  "PYTEST",        // Sunucu başlığı ayarla
		AppName:       "PYTEST v1.0.1", // Uygulama adını ayarla
	})
	app.Use(helmet.New(ConfigDefault))
	app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("CORS_URL"),
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true,
	}))

	//routing
	routes.Setup(app)
	app.Listen(":8080")

	
}


