package usecase

import (
	"context"
	"envmo/module/env_info/data/repository"
	"envmo/module/env_info/domain/entity"
)

type envInfoEmqxUsecaseImpl struct {
	envInfoRepository repository.EnvInfoEmqxRepository
}

// GetHistoryByDeviceID implements EnvInfoEmqxUsecase.
func (e *envInfoEmqxUsecaseImpl) GetHistoryByDeviceID(ctx context.Context, deviceID string, startTime int64, endTime int64) ([]entity.HistoryEnvInfoEntity, error) {
	return e.envInfoRepository.GetHistoryByDeviceID(ctx, deviceID,  startTime, endTime)
}

// ScheduleInsertHistoryEnvInfo implements EnvInfoEmqxUsecase.
func (e *envInfoEmqxUsecaseImpl) ScheduleInsertHistoryEnvInfo() {
	e.envInfoRepository.ScheduleInsertHistoryEnvInfo()
}

// Subscribe implements EnvInfoEmqxUsecase.
func (e *envInfoEmqxUsecaseImpl) Subscribe() {
	e.envInfoRepository.Subscribe()
}

// FakePublishEnvInfo implements EnvInfoEmqxDataSource.
func (e *envInfoEmqxUsecaseImpl) FakePublishEnvInfo() {
	e.envInfoRepository.FakePublish()
}

func NewEnvInfoEmqxUsecase(envInfoRepository repository.EnvInfoEmqxRepository) EnvInfoEmqxUsecase {

	return &envInfoEmqxUsecaseImpl{envInfoRepository: envInfoRepository}
}
