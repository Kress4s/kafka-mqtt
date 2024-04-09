package customer

import (
	"context"
	"fmt"
	"kafkaMq/config"
	kafka_tool "kafkaMq/kafkaTool"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type Customer struct {
	ctx    context.Context
	Reader *kafka.Reader
	Addr   string
	Port   string
}

func GetCustomer() *Customer {
	cfg := config.GetConfigure()
	return &Customer{
		ctx:    context.Background(),
		Reader: kafka_tool.GetKafkaReadConn(cfg),
	}
}

func (c *Customer) Read() {
	defer func() {
		if err := c.Reader.Close(); err != nil {
			panic(err)
		}
	}()
	for {
		log.Print("starting read...")
		m, err := c.Reader.ReadMessage(c.ctx)
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset,
			string(m.Key), string(m.Value))
		time.Sleep(1 * time.Second)
	}
}
