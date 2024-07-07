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

}
