package usecase

import (
	"envmo/module/env_info/data/repository"
)

type envInfoEmqxUsecaseImpl struct {
	envInfoRepository repository.EnvInfoEmqxRepository
}

// FakePublishEnvInfo implements EnvInfoEmqxDataSource.
func (e *envInfoEmqxUsecaseImpl) FakePublishEnvInfo() {
	e.envInfoRepository.FakePublishEnvInfo()
}

func NewEnvInfoEmqxUsecase(envInfoRepository repository.EnvInfoEmqxRepository) EnvInfoEmqxUsecase {

	return &envInfoEmqxUsecaseImpl{envInfoRepository: envInfoRepository}
}
