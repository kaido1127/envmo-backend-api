package infras

import (
	"context"
	"fmt"

	"com.pegatech.faceswap/app_config"
	"com.pegatech.faceswap/infras/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3ClientProvider struct {
	S3ClientInstance *s3.Client
	Bucket           *string
}

/*
GlobalMediaStorageRootPath = CdnEndPointUrl/Bucket

	Example: https://faceswap-staging.sgp1.cdn.digitaloceanspaces.com/faceswap-statging
*/
var GlobalMediaStorageRootPath *string

func CreateS3ClientProviderByConfig(awsS3Config app_config.AwsS3Config) *S3ClientProvider {

	s3Client := createS3Client(awsS3Config)

	if s3Client == nil {
		logger.Fatal("Failed to connect to S3 Client")
	}

	globalMediaStorageRootPathValue := fmt.Sprintf("%s/%s", awsS3Config.CdnEndpointUrl, awsS3Config.BucketName)
	GlobalMediaStorageRootPath = &globalMediaStorageRootPathValue

	return &S3ClientProvider{
		S3ClientInstance: s3Client,
		Bucket:           aws.String(awsS3Config.BucketName),
	}
}

func createS3Client(awsS3Config app_config.AwsS3Config) *s3.Client {
	endpoint := aws.Endpoint{
		URL:               awsS3Config.EndpointUrl,
		HostnameImmutable: true,
		Source:            aws.EndpointSourceCustom,
	}
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return endpoint, nil
	})

	sdkConfig, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsS3Config.AwsAccessKeyId, awsS3Config.AwsSecretAccessKeyId, "")),
		config.WithRegion(awsS3Config.RegionName),
	)

	if err != nil {
		logger.Fatal("Failed to load config for S3 Client", err)
	}

	s3Client := s3.NewFromConfig(sdkConfig)

	return s3Client
}
