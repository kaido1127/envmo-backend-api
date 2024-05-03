package infras

import (
	"com.pegatech.faceswap/app_config"
	"github.com/OneSignal/onesignal-go-api"
)

type OnesignalClientProvider struct {
	ApiClient *onesignal.APIClient
	AppId string
	RestApiKey string
} 

func CreateNewOnesignalClientByConfig(onesignalConfig app_config.OnesignalConfig) OnesignalClientProvider {

	configuration := onesignal.NewConfiguration()
	configuration.Debug = true
	apiClient := onesignal.NewAPIClient(configuration)

	return OnesignalClientProvider{
		ApiClient: apiClient,
		AppId:     onesignalConfig.AppId,
		RestApiKey: onesignalConfig.RestApiKey,
	}

}

