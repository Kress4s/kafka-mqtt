package mqtt_producer

import mqtt "github.com/eclipse/paho.mqtt.golang"

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

func PublishMessage(c mqtt.Client, msg interface{}, topic string, qos int, retained bool) bool {
	token := c.Publish(topic, byte(qos), retained, msg)
	return token.Wait()
}
