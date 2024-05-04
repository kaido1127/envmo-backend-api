package datasource

import (
	"encoding/json"
	infras "envmo/infras/emqx"
	"envmo/infras/logger"
	"envmo/module/env_info/data/model"
	"time"

	"golang.org/x/exp/rand"
)

type envInfoEmqxDataSourceImpl struct {
	emqxClient infras.EmqxClientProvider
}

// FakePublishEnvInfo implements EnvInfoEmqxDataSource.
func (e *envInfoEmqxDataSourceImpl) FakePublishEnvInfo() {
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
			token := e.emqxClient.Client.Publish(e.emqxClient.QueueName, 1, true, payload)
			//logger.Info("FakePublishEnvInfo","FinishPublish", )
			token.Wait()
			time.Sleep(10 * time.Second)
		}
	}()
}

func NewEnvInfoEmqxDataSource(emqxClient infras.EmqxClientProvider) EnvInfoEmqxDataSource {

	return &envInfoEmqxDataSourceImpl{emqxClient: emqxClient}
}

func fakeEnvInfoModel() model.EnvInfoModel {
	// Tạo một số giá trị mô phỏng cho độ ẩm và nhiệt độ
	humidity := rand.Float64() * 0.7    // Độ ẩm dao động từ 0 đến 70%
	temperatureInC := fakeTemperature() // Nhiệt độ được tạo mô phỏng

	// Tạo một vị trí ngẫu nhiên ở Việt Nam
	location := model.Location{
		Latitude:  getRandomLatitude(),
		Longitude: getRandomLongitude(),
	}

	// Trả về một mẫu dữ liệu mô phỏng
	return model.EnvInfoModel{
		CreatedAt:      time.Now().Unix(),
		Humidity:       humidity,
		TemperatureInC: temperatureInC,
		Location:       &location,
	}
}

func fakeTemperature() float64 {
	// Giả lập nhiệt độ dựa trên điều kiện thời tiết ở Việt Nam
	// Trong trường hợp này, chúng ta chỉ sử dụng các giá trị mô phỏng ngẫu nhiên
	return rand.Float64()*20 + 20 // Nhiệt độ dao động từ 20 đến 40 độ C
}

func getRandomLatitude() float64 {
	// Lấy một giá trị ngẫu nhiên cho vĩ độ trong phạm vi của Việt Nam
	// 8.18 đến 23.39 (phạm vi đất liền Việt Nam)
	return rand.Float64()*(23.39-8.18) + 8.18
}

func getRandomLongitude() float64 {
	// Lấy một giá trị ngẫu nhiên cho kinh độ trong phạm vi của Việt Nam
	// 102.14 đến 109.46 (phạm vi đất liền Việt Nam)
	return rand.Float64()*(109.46-102.14) + 102.14
}
