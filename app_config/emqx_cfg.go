package app_config

type EmqxConfig struct {
	Server            string `mapstructure:"Server"`
	ClientID          string `mapstructure:"ClientID"`
	QueueName         string `mapstructure:"QueueName"`
	PublisherName     string `mapstructure:"PublisherName"`
	PublisherPassword string `mapstructure:"PublisherPassword"`
}
