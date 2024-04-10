package mqtt_producer

import (
	"fmt"
	"kafkaMq/config"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type PublishOpts struct {
	Broker   string
	ClientID string
	UserName string
	Password string
}

func (p *PublishOpts) SetBroker(broker string) *PublishOpts {
	p.Broker = broker
	return p
}

func (p *PublishOpts) SetClientID(clientID string) *PublishOpts {
	p.ClientID = clientID
	return p
}

func (p *PublishOpts) SetOauth(username, password string) *PublishOpts {
	p.UserName = username
	p.Password = password
	return p
}

// func

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("接收topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
}

func GetPublisher(cfg *config.Configure) mqtt.Client {
	// opt := PublishOpts{}
	// opt.SetBroker(strings.Join([]string{cfg.Proto.TCP, cfg.Mqtt.Addr}, "//"))
	// opt.SetClientID("123")
	// opt.SetOauth(cfg.Mqtt.UserName, cfg.Mqtt.Password)
	// opt := mqtt.NewClientOptions().AddBroker(strings.Join([]string{cfg.Proto.TCP, cfg.Mqtt.Addr}, "//"))
	// opt := mqtt.NewClientOptions().AddBroker(strings.Join([]string{cfg.Proto.TCP, cfg.Mqtt.Addr}, "//"))
	opt := mqtt.NewClientOptions().AddBroker("tcp://121.41.38.13:1883")
	opt.SetClientID("123")
	opt.SetUsername(cfg.Mqtt.UserName)
	opt.SetPassword(cfg.Mqtt.Password)
	opt.CleanSession = false
	opt.SetDefaultPublishHandler(onMessageReceived)
	return mqtt.NewClient(opt)
}
