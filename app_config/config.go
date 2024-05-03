package app_config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

const (
	configFolder = "./app_config"
)

type AppConfig struct {
	ServerConfig    ServerConfig    `mapstructure:"server"`
	AwsS3Config     AwsS3Config     `mapstructure:"aws_s3"`
	FirebaseConfig  FirebaseConfig  `mapstructure:"firebase"`
	LoggerConfig    LoggerConfig    `mapstructure:"logger"`
	MongoDbConfig   MongoDbConfig   `mapstructure:"mongodb"`
	RabbitMqConfig  RabbitMQConfig  `mapstructure:"rabbitmq"`
	OnesignalConfig OnesignalConfig `mapstructure:"onesignal"`
}


func LoadAppConfig(configFileName string) AppConfig {
	configFilePath := fmt.Sprintf("%s/%s", configFolder, configFileName)
	viper.SetConfigFile(configFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to read config file", err)
	}
	appConfig := AppConfig{}
	err = viper.Unmarshal(&appConfig)
	if err != nil {
		log.Fatal("Failed to load config", err)
	}

	return appConfig
}
