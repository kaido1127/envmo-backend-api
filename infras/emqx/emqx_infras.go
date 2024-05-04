package infras

import (
	"log"
	"time"

	"envmo/app_config"
	"envmo/infras/logger"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	logger.Info("EMQX Queue Received ", "TOPIC", msg.Topic())
	logger.Info("EMQX Queue Received ", "MSG", msg.Payload())
}

type EmqxClientProvider struct {
	Client    mqtt.Client
	QueueName string
}

func CreateEmqxClientByConfig(config app_config.EmqxConfig) EmqxClientProvider {
	opts := mqtt.NewClientOptions().AddBroker(config.Server).SetClientID(config.ClientID)

	opts.SetKeepAlive(60 * time.Second)
	opts.SetUsername(config.PublisherName)
	opts.SetPassword(config.PublisherPassword)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	opts.OnConnect = func(c mqtt.Client) {
		log.Default().Println("EMQX Queue", "Connect", "Success")
	}
	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		log.Fatal("EMQX Queue", "ConnectErr", err)
	}
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("Emqx Infras", "Error", token.Error())
	}

	// Subscribe to a topic
	if token := c.Subscribe(config.QueueName, 0, nil); token.Wait() && token.Error() != nil {
		log.Fatal("EMQX Queue", "Subscribe "+config.QueueName+" Failed", token.Error())

	}

	return EmqxClientProvider{
		Client:    c,
		QueueName: config.QueueName,
	}

}
