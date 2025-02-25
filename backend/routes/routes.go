package routes

import (
	"github.com/gofiber/fiber/v2"
	//User
	servis "github.com/while/payproje/controllers/Servis"
	usercontroller "github.com/while/payproje/controllers/User"

	//Secure
	securecontroller "github.com/while/payproje/Secure"

	middlewares "github.com/while/payproje/midllewares"

	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Setup(app *fiber.App) {
	api := app.Group("/oauth/")

	app.Get("/metrics",securecontroller.RateLimiter(5), middlewares.UserMiddleTokenControl, monitor.New(monitor.Config{
		Title:      "Sistem Durumu",
		FontURL:    "https://fonts.googleapis.com/css2?family=Roboto:wght@400;900&display=swap%22",
		CustomHead: "Analitik Veriler",
		APIOnly:    false,
	}))

	//Servis (Satıcı) Routes
	api.Post("servis/addnumber", middlewares.UserMiddleTokenControlLevel(3), servis.AddNumer)       // OTP Oluştur
	api.Get("servis/getnumber", middlewares.UserMiddleTokenControlLevel(3), servis.NumaraGetir)    // OTP Oluştur
	api.Post("servis/wplogin", middlewares.UserMiddleTokenControlLevel(3), servis.LoginHandler) // OTP Oluştur
	api.Post("servis/wpstop", middlewares.UserMiddleTokenControlLevel(3), servis.StopHandler) // OTP Oluştur
	api.Post("servis/qr", middlewares.UserMiddleTokenControlLevel(3), servis.GetQRCodeHandler) // OTP Oluştur
	api.Post("servis/testmessage", middlewares.UserMiddleTokenControlLevel(3), servis.SendMessageHandler) // OTP Oluştur
	api.Post("servis/bot-status", middlewares.UserMiddleTokenControlLevel(3), servis.BotStatusHandler) // OTP Oluştur
	api.Get("servis/numbot/:id", middlewares.UserMiddleTokenControlLevel(3), servis.NumBot) // OTP Oluştur
	api.Get("servis/numbot-status", middlewares.UserMiddleTokenControlLevel(3), servis.StatusBot) // OTP Oluştur
	
	//User (Admin) Controller
	api.Post("user/create", usercontroller.CreateUser)                                                 // User Oluştur
	api.Post("user/otp", usercontroller.CreateUserOTP)                                                 // OTP Oluştur
	api.Post("user/login", usercontroller.Login)                                                       // User Login
	api.Post("user/verify", usercontroller.LoginVerify)                                                // OTP Login Verify
	api.Get("user/token", middlewares.UserMiddleTokenControl, usercontroller.TokenConrol)              // Token Control
	api.Post("user/action", middlewares.UserMiddleTokenControlLevel(5), usercontroller.UserFullUpdate) // User Güncelleme (Silme, Güncelleme, Aktiflik)"
	api.Get("user/get-all", middlewares.UserMiddleTokenControlLevel(3), usercontroller.GetFullUser)    // User Listesini getir
	api.Get("user/logout", usercontroller.UserLogout)                                                  // Çıkış yap
}
