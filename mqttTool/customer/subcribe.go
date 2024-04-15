package mqtt_customer

import (
	"fmt"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// var OnMessageReceivedMapClientID map[string]mqtt.MessageHandler

// this function is deal with message when this client subscribed mqtt topic that be published message
func OnMessageReceived(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("接收topic: %s\nMessage: %s\n", message.Topic(), message.Payload())

}

// func GetMessageReceivedHandler(ClientID string) mqtt.MessageHandler

func Subscriber(c mqtt.Client, topic string, qos int) {
	// OnMessageReceived handler define in here, can flexible
	if token := c.Subscribe(topic, byte(qos), OnMessageReceived); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
		os.Exit(1)
	}
}
