package app_config

type EmqxConfig struct {
	Server             string `mapstructure:"Server"`
	PubClientID        string `mapstructure:"PubClientID"`
	SubClientID        string `mapstructure:"SubClientID"`
	QueueName          string `mapstructure:"QueueName"`
	PublisherName      string `mapstructure:"PublisherName"`
	PublisherPassword  string `mapstructure:"PublisherPassword"`
	SubscriberName     string `mapstructure:"SubscriberName"`
	SubscriberPassword string `mapstructure:"SubscriberPassword"`
}
