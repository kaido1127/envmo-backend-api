package envinfo

import (
	infras "envmo/infras/emqx"
	datasource "envmo/module/env_info/data/data_source"
	"envmo/module/env_info/data/repository"
	"envmo/module/env_info/domain/usecase"
)

func SetupEnvInfo(emqxClient infras.EmqxClientProvider) {
	envInfoDataSource := datasource.NewEnvInfoEmqxDataSource(emqxClient)
	envInfoRepository := repository.NewEnvInfoEmqxRepository(envInfoDataSource)
	envInfoUsecase := usecase.NewEnvInfoEmqxUsecase(envInfoRepository)

	go func() {
		envInfoUsecase.FakePublishEnvInfo()
	}()
}
