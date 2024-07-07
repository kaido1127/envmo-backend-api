package main

import (

	// aws_infras "envmo/infras/aws"
	// firebase_infras "envmo/infras/firebase"
	// mongodb_infras "envmo/infras/mongodb"
	// onesignal_infras "envmo/infras/onesignal"
	// rabbitmq_infras "envmo/infras/rabbitmq"

	"envmo/app_config"
	emqx_infras "envmo/infras/emqx"
	firebase_infras "envmo/infras/firebase"
	zap_logger "envmo/infras/logger"
	mongo_infras "envmo/infras/mongodb"
	envinfo "envmo/module/env_info"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	configApp := app_config.LoadAppConfig()

	// Logger init by config
	initLoggerErr := zap_logger.CreateLoggerByConfig(configApp.LoggerConfig)
	if initLoggerErr != nil {
		log.Fatal("Failed to create ZapLogger", initLoggerErr)
	}
	zap_logger.Info("Complete ZapLogger configuration")
	defer zap_logger.Sync()

	// MongoDB init by config
	mongoDB := mongo_infras.CreateMongoDBProviderByConfig(configApp.MongoDbConfig)
	zap_logger.Info("Complete MongoDB configuration")

	// // AWS S3 init by config
	// s3Client := aws_infras.CreateS3ClientProviderByConfig(configApp.AwsS3Config)
	// zap_logger.Info("Complete AWS S3 configuration")

	// // RabbitMQ init by config
	// rabbitMQ := rabbitmq_infras.CreateRabbitMQProviderByConfig(configApp.RabbitMqConfig)
	// zap_logger.Info("Complete RabbitMQ configuration")

	// // Onesignal init by config
	// onesignal := onesignal_infras.CreateNewOnesignalClientByConfig(configApp.OnesignalConfig)
	// zap_logger.Info("Complete Onesignal configuration")
	// FirebaseAuth config

	// Emqx init by config
	emqxPubClient := emqx_infras.CreateEmqxPublisherByConfig(configApp.EmqxConfig)
	zap_logger.Info("Complete EmqxPub configuration")

	emqxSubClient := emqx_infras.CreateEmqxSubscriberByConfig(configApp.EmqxConfig)
	zap_logger.Info("Complete EmqxPub configuration")

	// RealtimeDB init by config
	realtimeDbClient := firebase_infras.CreateRealtimeDbClientByConfig(configApp.FirebaseConfig)
	zap_logger.Info("Complete Firebase configuration")

	app := fiber.New()
	app.Use(logger.New(
		logger.Config{
			Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		},
	))

	envinfo.SetupEnvInfo(app,emqxPubClient, emqxSubClient, *realtimeDbClient, *mongoDB)

	//zap_logger.Info("Complete FirebaseAuth configuration")

	// video_template.SetupVideoTemplate(app, mongoDB)
	// image_template.SetupImageTemplate(app, mongoDB)
	// userprofile.SetupUserProfile(app, s3Client, mongoDB)
	// storage_media.SetupStorageMedia(app, s3Client, mongoDB)
	// task.SetupTask(app, mongoDB, rabbitMQ, realtimeDbClient)

	// notification.NotificationService(&rabbitMQ, mongoDB, &onesignal)
	// task_processing.TaskProcessingService(&rabbitMQ, realtimeDbClient, mongoDB)

	appErr := app.Listen(configApp.ServerConfig.HttpPort)
	if appErr != nil {
		zap_logger.Fatal("Failed to start server")
	}

}
