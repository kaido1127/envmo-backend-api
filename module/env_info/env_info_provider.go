package envinfo

import (
	emqx_infras "envmo/infras/emqx"
	"envmo/module/env_info/api/controller"
	"envmo/module/env_info/api/router"
	datasource "envmo/module/env_info/data/data_source"
	"envmo/module/env_info/data/repository"
	"envmo/module/env_info/domain/usecase"

	"firebase.google.com/go/v4/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupEnvInfo(app *fiber.App, emqxPubClient emqx_infras.EmqxClientProvider, emqxSubClient emqx_infras.EmqxClientProvider, realtimeDB db.Client, mongoDb mongo.Database) {
	envInfoEmqxDataSource := datasource.NewEnvInfoEmqxDataSource(emqxPubClient, emqxSubClient)
	envInfoRealtimeDataSource := datasource.NewEnvInfoRealtimeDataSource(realtimeDB)
	envInfoMongoDataSource := datasource.NewEnvInfoMongoDataSource(mongoDb)
	envInfoRepository := repository.NewEnvInfoEmqxRepository(envInfoEmqxDataSource, envInfoRealtimeDataSource, envInfoMongoDataSource)
	envInfoUsecase := usecase.NewEnvInfoEmqxUsecase(envInfoRepository)
	envInfoController := controller.NewEnvController(envInfoUsecase)
	
	go func() {
		envInfoUsecase.FakePublishEnvInfo()
	}()
	go func() {
		envInfoUsecase.Subscribe()
	}()
	
	envInfoUsecase.ScheduleInsertHistoryEnvInfo()
	
	router.SetUpRoutes(app, envInfoController)
}
