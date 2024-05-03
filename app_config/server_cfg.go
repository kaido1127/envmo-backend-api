package app_config

type ServerConfig struct {
	HttpPort string `mapstructure:"HttpPort"`
	Mode     string `mapstructure:"Mode"`
}


