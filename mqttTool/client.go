package mqtt_tool

import (
	"kafkaMq/config"
	"log"
	"os"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// different client has itself client
var clientMap = make(map[string]mqtt.Client, 0)

func GetClient(cfg *config.Configure, clientID string) mqtt.Client {
	if _, ok := clientMap[clientID]; !ok {
		opt := mqtt.NewClientOptions().AddBroker(strings.Join([]string{cfg.Proto.TCP, cfg.Mqtt.Addr}, "//"))
		opt.SetClientID(clientID)
		opt.SetUsername(cfg.Mqtt.UserName)
		opt.SetPassword(cfg.Mqtt.Password)
		// opt.SetCleanSession(false)
		opt.CleanSession = false
		// opt.SetDefaultPublishHandler(customer.OnMessageReceived)
		client := mqtt.NewClient(opt)
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			log.Fatal(token.Error())
			os.Exit(1)
		}
		clientMap[clientID] = client
		return client
	}
	return clientMap[clientID]
}

func DisClientConnect(client mqtt.Client, quiesce uint) {
	client.Disconnect(quiesce)
}
