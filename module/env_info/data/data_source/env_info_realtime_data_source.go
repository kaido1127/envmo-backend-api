package datasource

import (
	"context"
	"envmo/module/env_info/data/model"
)

type EnvInfoRealtimeDataSource interface {
	Update(ctx context.Context, userEnvInfoModel model.UserEnvInfoModel) error
}