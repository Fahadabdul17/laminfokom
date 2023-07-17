package controller

import (
	"github.com/aiteung/musik"
	"github.com/fahadabdul17/hadbackend"
	"github.com/fahadabdul17/laminfokom/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/whatsauth/whatsauth"
)

var DataAkreditas = "dataakreditas"
var DataProgramStudi = "dataprogramstudi"
var Profile = "profile"

func WsWhatsAuthQR(c *websocket.Conn) {
	whatsauth.RunSocket(c, config.PublicKey, config.Usertables[:], config.Ulbimariaconn)
}

func PostWhatsAuthRequest(c *fiber.Ctx) error {
	if string(c.Request().Host()) == config.Internalhost {
		var req whatsauth.WhatsauthRequest
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}
		ntfbtn := whatsauth.RunModuleLegacy(req, config.PrivateKey, config.Usertables[:], config.Ulbimariaconn)
		return c.JSON(ntfbtn)
	} else {
		var ws whatsauth.WhatsauthStatus
		ws.Status = string(c.Request().Host())
		return c.JSON(ws)
	}

}

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetDataAkreditas(c *fiber.Ctx) error {
	getstatus := hadbackend.GetDataAkreditas("Masih Berlaku")
	return c.JSON(getstatus)
}

func GetDataProgramStudi(c *fiber.Ctx) error {
	getstatus := hadbackend.GetDataProgramStudi("Sarjana Terapan")
	return c.JSON(getstatus)
}

func GetProfile(c *fiber.Ctx) error {
	getstatus := hadbackend.GetDataProfile("kampus merdeka")
	return c.JSON(getstatus)
}
