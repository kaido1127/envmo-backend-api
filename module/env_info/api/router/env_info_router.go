package router

import (
	"envmo/module/env_info/api/controller"

	"github.com/gofiber/fiber/v2"
)

const (
	EnvmoInfoGroup = "/api/env-info"
)

// http://127.0.0.1:3000/api/env-info/history?device_id=fake_device_id_1&page=1&limit=100
func SetUpRoutes(app *fiber.App, controller *controller.EnvmoController) {

	envInfo := app.Group(EnvmoInfoGroup)
	envInfo.Get("/history", controller.GetEnvmoHistoryHandler)
	envInfo.Post("/devices/14-18-C3-3B-A4-8E/subscribe", func(c *fiber.Ctx) error {
		return c.SendString("Subscribed 14-18-C3-3B-A4-8E successfully")
	})
	envInfo.Post("/devices/14-18-C3-3B-A4-8E/unsubscribe", func(c *fiber.Ctx) error {
		return c.SendString("Unsubscribed 14-18-C3-3B-A4-8E successfully")
	})
}
