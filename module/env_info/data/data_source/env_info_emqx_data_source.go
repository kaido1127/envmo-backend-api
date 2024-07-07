package datasource

import "envmo/module/env_info/data/model"

type EnvInfoEmqxDataSource interface {
	FakePublish()
	Subscribe(processEnvInfoCallback func(envInfoModel model.EnvInfoModel) error) // Subscribe
}
