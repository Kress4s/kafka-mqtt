package producer

import (
	"context"
	"kafkaMq/config"
	kafka_tool "kafkaMq/kafkaTool"
	"sync"

	"github.com/segmentio/kafka-go"
)

/*
kafka producer:
	this file will generate message to kafka broker  to test kafka's function and performance
*/

var (
	once     sync.Once
	producer Producer
)

type Producer struct {
	ctx    context.Context
	Writer *kafka.Writer
	Addr   string
	Port   string
}

func GetProducer() *Producer {
	cfg := config.GetConfigure()
	once.Do(func() {
		producer = Producer{
			ctx:    context.Background(),
			Writer: kafka_tool.GetKafkaWriteConn(cfg),
			Addr:   cfg.Kafka.Addr,
			Port:   cfg.Kafka.Port,
		}
	})
	return &producer
}

func (p *Producer) Send(msg ...kafka.Message) error {
	err := p.Writer.WriteMessages(p.ctx, msg...)
	defer p.Writer.Close()
	if err != nil {
		return err
	}
	return nil
}
