package controller

import (
	"envmo/common/errors/error_response"
	validator_module "envmo/module/env_info/api/validator"
	"envmo/module/env_info/domain/usecase"

	//"envmo/module/video_template/api/middleware"

	"github.com/gofiber/fiber/v2"
)

type EnvmoController struct {
	envmoUsecase usecase.EnvInfoEmqxUsecase
}

func NewEnvController(envmoUsecase usecase.EnvInfoEmqxUsecase) *EnvmoController {
	return &EnvmoController{envmoUsecase: envmoUsecase}
}

func (c *EnvmoController) GetEnvmoHistoryHandler(ctx *fiber.Ctx) error {

	historyQuery := new(validator_module.HistoryQuery)
	if err := ctx.QueryParser(historyQuery); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(error_response.BadRequestErrorResponse("Bad request"))
	}

	validator := validator_module.NewHistoryQueryValidator()
	if err := validator.Validate(historyQuery); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	deviceID := historyQuery.DeviceID
	startTime := historyQuery.StartTime
	endTime := historyQuery.EndTime

	history, err := c.envmoUsecase.GetHistoryByDeviceID(ctx.Context(), deviceID, startTime, endTime)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(error_response.InternalServerErrorResponse("Failed to get history of this device"))
	}

	return ctx.JSON(history)
}
