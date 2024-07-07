package repository

import (
	"context"
	"envmo/infras/logger"
	datasource "envmo/module/env_info/data/data_source"
	"envmo/module/env_info/data/model"
	"envmo/module/env_info/domain/entity"
	"envmo/module/env_info/dto"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

type envInfoEmqxRepositoryImpl struct {
	envInfoEmqxDataSource     datasource.EnvInfoEmqxDataSource
	envInfoRealtimeDataSource datasource.EnvInfoRealtimeDataSource
	envInfoMongoDataSource    datasource.EnvInfoMongoDataSource
}

// GetHistoryByDeviceID implements EnvInfoEmqxRepository.
func (e *envInfoEmqxRepositoryImpl) GetHistoryByDeviceID(ctx context.Context, deviceID string, startTime int64, endTime int64) ([]entity.HistoryEnvInfoEntity, error) {
	models, err := e.envInfoMongoDataSource.GetHistoryByDeviceID(ctx, deviceID, startTime, endTime)
	if err != nil {
		return nil, err
	}
	return dto.HistoryEnvInfoEntitiesFromModels(models), nil
}

const scheduleTimeInMinutes = 15

var latestHistoryEnvInfoMap = make(map[string]model.HistoryEnvInfoModel)

// ScheduleInsertHistoryEnvInfo implements EnvInfoEmqxRepository.
func (e *envInfoEmqxRepositoryImpl) ScheduleInsertHistoryEnvInfo() {
	c := cron.New()

	spec := fmt.Sprintf("*/%d * * * *", scheduleTimeInMinutes)
	c.AddFunc(spec, func() {
		if len(latestHistoryEnvInfoMap) > 0 {
			var validHistoryEnvInfoModels []model.HistoryEnvInfoModel

			for _, info := range latestHistoryEnvInfoMap {
				validHistoryEnvInfoModels = append(validHistoryEnvInfoModels, info)
			}
			count := len(validHistoryEnvInfoModels)
			logger.Info("EnvInfoEmqxRepository",
				"ScheduleInsertHistoryEnvInfo", fmt.Sprintf("Start insert new %d records", count),
			)
			updatedCount, err := e.envInfoMongoDataSource.InsertMany(context.Background(), validHistoryEnvInfoModels)
			if err != nil {
				logger.Error("EnvInfoEmqxRepository",
					"ScheduleInsertHistoryEnvInfo", err,
				)
			} else {
				logger.Info("EnvInfoEmqxRepository",
					"ScheduleInsertHistoryEnvInfo", "Finish instert, continue deleting "+fmt.Sprintf("%d", scheduleTimeInMinutes)+" minutes later",
					"Ratio update", fmt.Sprintf("%d/%d", updatedCount, count),
				)
			}
		}
	})

	c.Start()
}

func (e *envInfoEmqxRepositoryImpl) processEnvInfoCallback(envInfoModel model.EnvInfoModel) error {
	//Đảm bảo record mới phải tạo trong vòng scheduleTimeInMinutes gần đây, nếu không sẽ insert 1 bản record cũ
	if envInfoModel.CreatedAt != nil && *envInfoModel.CreatedAt >= (time.Now().Unix()-scheduleTimeInMinutes*60) {
		latestHistoryEnvInfoMap[envInfoModel.MacAddr] = dto.HistoryEnvInfoModelFromQueueModel(envInfoModel)
	}
	return e.envInfoRealtimeDataSource.Update(context.Background(), dto.UserEnvInfoModelFromQueueModel(envInfoModel))
}

// Subscribe implements EnvInfoEmqxRepository.
func (e *envInfoEmqxRepositoryImpl) Subscribe() {
	e.envInfoEmqxDataSource.Subscribe(e.processEnvInfoCallback)
}

// FakePublish implements EnvInfoEmqxDataSource.
func (e *envInfoEmqxRepositoryImpl) FakePublish() {
	e.envInfoEmqxDataSource.FakePublish()
}

func NewEnvInfoEmqxRepository(envInfoEmqxDataSource datasource.EnvInfoEmqxDataSource,
	envInfoRealtimeDataSource datasource.EnvInfoRealtimeDataSource,
	envInfoMongoDataSource datasource.EnvInfoMongoDataSource) EnvInfoEmqxRepository {

	return &envInfoEmqxRepositoryImpl{
		envInfoEmqxDataSource:     envInfoEmqxDataSource,
		envInfoRealtimeDataSource: envInfoRealtimeDataSource,
		envInfoMongoDataSource:    envInfoMongoDataSource,
	}
}
