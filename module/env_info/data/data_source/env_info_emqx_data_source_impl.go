package datasource

import (
	"encoding/json"
	infras "envmo/infras/emqx"
	"envmo/infras/logger"
	"envmo/module/env_info/data/model"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"golang.org/x/exp/rand"
)

type envInfoEmqxDataSourceImpl struct {
	emqxPubClient infras.EmqxClientProvider
	emqxSubClient infras.EmqxClientProvider
}

// Subscribe implements EnvInfoEmqxDataSource.
func (e *envInfoEmqxDataSourceImpl) Subscribe(processEnvInfoCallback func(envInfoModel model.EnvInfoModel) error) {

	e.emqxSubClient.Options.SetDefaultPublishHandler(
		func(client mqtt.Client, msg mqtt.Message) {
			msg.Ack()
			logger.Info("EMQX Queue Received ", "TOPIC", msg.Topic())
			logger.Info("EMQX Queue Received ", "MSG", msg.Payload())

			var envInfoModel *model.EnvInfoModel
			err := json.Unmarshal(msg.Payload(), &envInfoModel)
			if err != nil {
				logger.Error("EnvInfoEmqxDataSource",
					"DecodeDataError", err,
					"MessageBody", msg.Payload())
			} else {
				logger.Info("EnvInfoEmqxDataSource",
					"MessageBody", envInfoModel)
			}

			callbackErr := processEnvInfoCallback(*envInfoModel)
			if callbackErr == nil {
				logger.Info("EnvInfoEmqxDataSource",
					"CallbackSloveEnvInfoModelSuccess", envInfoModel)
			} else {
				logger.Error("EnvInfoEmqxDataSource",
					"CallbackSloveEnvInfoModelError",
					envInfoModel, "Error", callbackErr)
			}
		},
	)
	client := mqtt.NewClient(e.emqxSubClient.Options)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logger.Error("Emqx Infras", "Error", token.Error())
	}

	if token := client.Subscribe(e.emqxSubClient.QueueName, 0, nil); token.Wait() && token.Error() != nil {
		logger.Error("EMQX Queue", "Subscribe "+e.emqxSubClient.QueueName+" Failed", token.Error())
	} else {
		logger.Info("EMQX Queue", "SubscribeSuccess ", e.emqxSubClient.QueueName)
	}
}

// FakePublish implements EnvInfoEmqxDataSource.
func (e *envInfoEmqxDataSourceImpl) FakePublish() {
	client := mqtt.NewClient(e.emqxPubClient.Options)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logger.Fatal("Emqx Infras", "Error", token.Error())
	}

	go func() {
		for {
			fakeEnvInfoModel := fakeEnvInfoModel()
			payload, err := json.Marshal(fakeEnvInfoModel)
			if err != nil {
				logger.Error("FakePublishEnvInfo", "JsonToPayloadErr", err)
			}
			logger.Info("FakePublishEnvInfo",
				"StartPublish", fakeEnvInfoModel,
				"Payload", payload)
			token := client.Publish(e.emqxPubClient.QueueName, 1, true, payload)
			//logger.Info("FakePublishEnvInfo","FinishPublish", )
			token.Wait()
			time.Sleep(3 * time.Second)
		}
	}()
}

func NewEnvInfoEmqxDataSource(emqxPubClient infras.EmqxClientProvider, emqxSubClient infras.EmqxClientProvider) EnvInfoEmqxDataSource {

	return &envInfoEmqxDataSourceImpl{emqxPubClient: emqxPubClient,
		emqxSubClient: emqxSubClient}
}

func fakeEnvInfoModel() model.EnvInfoModel {
	humidity := rand.Float32() * 70
	temperatureInC := randomTemperatureInC()
	pressureInHpa := randomPressureInHpa()
	gasInPpm := randomGasInPpm()

	location := "Unknown"
	createdAt := time.Now().Unix()

	return model.EnvInfoModel{
		MacAddr:           "14-18-C3-3B-A4-8E",
		TemperatureInC:    &temperatureInC,
		HumidityInPercent: &humidity,
		PressureInHpa:     &pressureInHpa,
		GasInPpm:          &gasInPpm,
		CreatedAt:         &createdAt,
		Location:          &location,
	}
}

func randomGasInPpm() float32 {
	rand.Seed(uint64(time.Now().UnixNano()))
	minGas := 600.0
	maxGas := 650.0
	return float32(minGas + rand.Float64()*(maxGas-minGas))
}

func randomPressureInHpa() float32 {
	rand.Seed(uint64(time.Now().UnixNano()))
	minPressure := 980.0
	maxPressure := 1050.0
	return float32(minPressure + rand.Float64()*(maxPressure-minPressure))
}

var baseTemperature float32 = 31.0

func randomTemperatureInC() float32 {
	rand.Seed(uint64(time.Now().UnixNano()))
	minVariation := -2.0
	maxVariation := 2.0
	variation := float32(minVariation + rand.Float64()*(maxVariation-minVariation))
	return baseTemperature + variation
}

// func getRandomLatitude() float64 {
// 	// Lấy một giá trị ngẫu nhiên cho vĩ độ trong phạm vi của Việt Nam
// 	// 8.18 đến 23.39 (phạm vi đất liền Việt Nam)
// 	return rand.Float64()*(23.39-8.18) + 8.18
// }

// func getRandomLongitude() float64 {
// 	// Lấy một giá trị ngẫu nhiên cho kinh độ trong phạm vi của Việt Nam
// 	// 102.14 đến 109.46 (phạm vi đất liền Việt Nam)
// 	return rand.Float64()*(109.46-102.14) + 102.14
// }
