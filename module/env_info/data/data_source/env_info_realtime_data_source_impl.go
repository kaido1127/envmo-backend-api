package datasource

import (
	"context"
	"envmo/module/env_info/data/model"
	"firebase.google.com/go/v4/db"
)

type envInfoRealtimeDataSourceImpl struct {
	realtimeDb db.Client
}

// Upsert implements EnvInfoRealtimeDataSource.
func (e envInfoRealtimeDataSourceImpl) Update(ctx context.Context, userEnvInfoModel model.UserEnvInfoModel) error {
	ref := e.realtimeDb.NewRef("devices/" + userEnvInfoModel.DeviceID)
	return ref.Set(ctx, userEnvInfoModel.ToMap())
}

func NewEnvInfoRealtimeDataSource(realtimeDb db.Client) EnvInfoRealtimeDataSource {
	return envInfoRealtimeDataSourceImpl{
		realtimeDb: realtimeDb,
	}
}
