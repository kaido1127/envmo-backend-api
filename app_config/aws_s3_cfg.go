package app_config


type AwsS3Config struct {
	EndpointUrl          string `mapstructure:"EndpointUrl"`
	CdnEndpointUrl       string `mapstructure:"CdnEndpointUrl"`
	AwsAccessKeyId       string `mapstructure:"AwsAccessKeyId"`
	AwsSecretAccessKeyId string `mapstructure:"AwsSecretAccessKeyId"`
	RegionName           string `mapstructure:"RegionName"`
	BucketName           string `mapstructure:"BucketName"`
}

