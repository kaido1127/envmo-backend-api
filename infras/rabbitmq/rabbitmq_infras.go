package infras

import (
	"envmo/infras/logger"

	"envmo/app_config"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQProvider struct {
	Channel            *amqp.Channel
	TaskQueueName      string
	CompletedTaskQueue string
}

func CreateRabbitMQProviderByConfig(rabbitMqConfig app_config.RabbitMQConfig) RabbitMQProvider {
	return RabbitMQProvider{
		Channel:            createRabbitMQChannel(rabbitMqConfig),
		TaskQueueName:      rabbitMqConfig.TaskQueue,
		CompletedTaskQueue: rabbitMqConfig.CompletedTaskQueue,
	}
}

func createRabbitMQChannel(rabbitMqConfig app_config.RabbitMQConfig) *amqp.Channel {

	connection, err := amqp.DialConfig(rabbitMqConfig.URI, amqp.Config{
		Heartbeat: 0,
	})

	if err != nil {
		logger.Fatal("Failed to load DialConfig RabbitMQ", err)
	}

	//defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		logger.Fatal("Failed to connect to RabbitMQ Channel", err)
	}

	_, err = channel.QueueDeclare(
		rabbitMqConfig.TaskQueue, // tên hàng đợi
		true,                     // durable
		false,                    // auto-delete
		false,                    // exclusive
		false,                    // no-wait
		amqp.Table{
			"x-max-priority": 10,
		},
	)
	if err != nil {
		logger.Fatal("Failed to Declare TaskQueue", err)
	}
	_, err = channel.QueueDeclare(
		rabbitMqConfig.CompletedTaskQueue, // tên hàng đợi
		true,                              // durable
		false,                             // auto-delete
		false,                             // exclusive
		false,                             // no-wait
		amqp.Table{
			"x-max-priority": 10,
		},
	)
	channel.Qos(1, 0, false)
	if err != nil {
		logger.Fatal("Failed to Declare CompletedTaskQueue", err)
	}
	return channel
}
