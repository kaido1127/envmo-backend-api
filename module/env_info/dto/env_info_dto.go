package dto

import (
	"envmo/module/env_info/data/model"
	"envmo/module/env_info/domain/entity"
	"math"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserEnvInfoModelFromQueueModel(envInfoModel model.EnvInfoModel) model.UserEnvInfoModel {
	var createdAt int64
	if envInfoModel.CreatedAt == nil {
		createdAt = time.Now().Unix()
	} else {
		createdAt = *envInfoModel.CreatedAt
	}
	return model.UserEnvInfoModel{
		DeviceID:          envInfoModel.MacAddr,
		TemperatureInC:    roundToNullable(envInfoModel.TemperatureInC),
		HumidityInPercent: roundToNullable(envInfoModel.HumidityInPercent),
		PressureInHpa:     roundToNullable(envInfoModel.PressureInHpa),
		GasInPpm:          roundToNullable(envInfoModel.GasInPpm),
		CreatedAt:         createdAt,
		Location:          envInfoModel.Location,
	}
}

func HistoryEnvInfoModelFromQueueModel(envInfoModel model.EnvInfoModel) model.HistoryEnvInfoModel {
	var createdAt int64
	if envInfoModel.CreatedAt == nil {
		createdAt = time.Now().Unix()
	} else {
		createdAt = *envInfoModel.CreatedAt
	}
	return model.HistoryEnvInfoModel{
		ID:                primitive.NewObjectID(),
		DeviceID:          envInfoModel.MacAddr,
		TemperatureInC:    roundToNullable(envInfoModel.TemperatureInC),
		PressureInHpa:     roundToNullable(envInfoModel.PressureInHpa),
		HumidityInPercent: roundToNullable(envInfoModel.HumidityInPercent),
		GasInPpm:          roundToNullable(envInfoModel.GasInPpm),
		CreatedAt:         createdAt,
		Location:          envInfoModel.Location,
	}
}

func HistoryEnvInfoEntityFromModel(envInfoModel model.HistoryEnvInfoModel) entity.HistoryEnvInfoEntity {
	return entity.HistoryEnvInfoEntity{
		ID:                envInfoModel.ID.Hex(),
		DeviceID:          envInfoModel.DeviceID,
		TemperatureInC:    roundToNullable(envInfoModel.TemperatureInC),
		PressureInHpa:     roundToNullable(envInfoModel.PressureInHpa),
		HumidityInPercent: roundToNullable(envInfoModel.HumidityInPercent),
		GasInPpm:          roundToNullable(envInfoModel.GasInPpm),
		CreatedAt:         envInfoModel.CreatedAt,
		Location:          envInfoModel.Location,
	}
}

func HistoryEnvInfoEntitiesFromModels(models []model.HistoryEnvInfoModel) []entity.HistoryEnvInfoEntity {
	lengthList := len(models)
	entities := make([]entity.HistoryEnvInfoEntity, lengthList)

	for i := 0; i < lengthList; i++ {
		entities[i] = HistoryEnvInfoEntityFromModel(models[i])
	}

	return entities
}

func HistoryEnvInfoModelFromUserModel(envInfoModel model.UserEnvInfoModel) model.HistoryEnvInfoModel {
	return model.HistoryEnvInfoModel{
		ID:                primitive.NewObjectID(),
		DeviceID:          envInfoModel.DeviceID,
		TemperatureInC:    roundToNullable(envInfoModel.TemperatureInC),
		PressureInHpa:     roundToNullable(envInfoModel.PressureInHpa),
		HumidityInPercent: roundToNullable(envInfoModel.HumidityInPercent),
		GasInPpm:          roundToNullable(envInfoModel.GasInPpm),
		CreatedAt:         envInfoModel.CreatedAt,
		Location:          envInfoModel.Location,
	}
}

func roundToTwoDecimalPlaces(value float32) float32 {
    return float32(math.Round(float64(value)*100) / 100)
}

// Helper function to round nullable float values
func roundToNullable(value *float32) *float32 {
	if value != nil {
		roundedValue := roundToTwoDecimalPlaces(*value)
		return &roundedValue
	}
	return nil
}
