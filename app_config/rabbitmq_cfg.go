package app_config

type RabbitMQConfig struct {
	URI                string `mapstructure:"URI"`
	TaskQueue          string `mapstructure:"TaskQueue"`
	CompletedTaskQueue string `mapstructure:"CompletedTaskQueue"`
}

