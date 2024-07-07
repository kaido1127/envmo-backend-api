package repository

import (
	"context"
	"envmo/module/env_info/domain/entity"
)

type EnvInfoEmqxRepository interface {
	FakePublish()
	Subscribe()
	ScheduleInsertHistoryEnvInfo()
	GetHistoryByDeviceID(ctx context.Context, deviceID string, startTime int64, endTime int64) ([]entity.HistoryEnvInfoEntity, error)
}
