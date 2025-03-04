package url

import (
	"github.com/fahadabdul17/laminfokom/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Homepage)              //API from user whatsapp message from iteung gowa
	page.Get("/dataakreditas", controller.GetDataAkreditas) //API from user whatsapp message from iteung gowa
	page.Get("/dataprogramstudi", controller.GetDataProgramStudi)   //API from user whatsapp message from iteung gowa
	page.Get("/profile", controller.GetProfile)   //API from user whatsapp message from iteung gowa
	
}