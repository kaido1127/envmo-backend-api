package app_config

type MongoDbConfig struct {
	ConnectionString string `mapstructure:"ConnectionString"`
	DbName           string `mapstructure:"DbName"`
}
