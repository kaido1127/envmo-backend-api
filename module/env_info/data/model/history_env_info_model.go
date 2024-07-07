package model

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	EnvInfoMongoCollectionName = "env_info"
)

type HistoryEnvInfoModel struct {
	ID                primitive.ObjectID `bson:"_id"`
	DeviceID          string             `bson:"device_id"`
	TemperatureInC    *float32           `bson:"temperature_in_c,omitempty"`
	HumidityInPercent *float32           `bson:"humidity_in_percent,omitempty"`
	PressureInHpa     *float32           `bson:"pressure_in_hpa,omitempty"`
	GasInPpm          *float32           `bson:"gas_in_ppm,omitempty"`
	CreatedAt         int64              `bson:"created_at"`
	Location          *string            `bson:"location,omitempty"`
}
