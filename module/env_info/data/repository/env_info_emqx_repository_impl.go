package repository

import datasource "envmo/module/env_info/data/data_source"

type envInfoEmqxRepositoryImpl struct {
	envInfoDataSource datasource.EnvInfoEmqxDataSource
}

// FakePublishEnvInfo implements EnvInfoEmqxDataSource.
func (e *envInfoEmqxRepositoryImpl) FakePublishEnvInfo() {
	e.envInfoDataSource.FakePublishEnvInfo()
}

func NewEnvInfoEmqxRepository(envInfoDataSource datasource.EnvInfoEmqxDataSource) EnvInfoEmqxRepository {

	return &envInfoEmqxRepositoryImpl{envInfoDataSource: envInfoDataSource}
}
