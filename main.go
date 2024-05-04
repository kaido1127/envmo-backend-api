package main

import (

	// aws_infras "envmo/infras/aws"
	// firebase_infras "envmo/infras/firebase"
	// mongodb_infras "envmo/infras/mongodb"
	// onesignal_infras "envmo/infras/onesignal"
	// rabbitmq_infras "envmo/infras/rabbitmq"

	"envmo/app_config"
	emqx_infras "envmo/infras/emqx"
	zap_logger "envmo/infras/logger"
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
	// mongoDB := mongodb_infras.CreateMongoDBProviderByConfig(configApp.MongoDbConfig)
	// zap_logger.Info("Complete MongoDB configuration")

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
	emqxClient := emqx_infras.CreateEmqxClientByConfig(configApp.EmqxConfig)
	zap_logger.Info("Complete Emqx configuration", "emqx", emqxClient.Client)

	envinfo.SetupEnvInfo(emqxClient)
	
	app := fiber.New()
	app.Use(logger.New(
		logger.Config{
			Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		},
	))

	//zap_logger.Info("Complete FirebaseAuth configuration")

	// RealtimeDB config
	//realtimeDbClient := firebase_infras.CreateRealtimeDbClientByConfig(configApp.FirebaseConfig)

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
