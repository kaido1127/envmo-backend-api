package datasource

import (
	"context"
	"envmo/module/env_info/data/model"
)

type EnvInfoMongoDataSource interface {
	InsertMany(ctx context.Context, historyEnvInfoModel []model.HistoryEnvInfoModel) (int, error)
	GetHistoryByDeviceID(ctx context.Context, deviceID string, startTime int64, endTime int64) ([]model.HistoryEnvInfoModel, error)
}
