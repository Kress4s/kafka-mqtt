package kafka_customer

import (
	"context"
	"database/sql"
	"fmt"
	"kafkaMq/config"
	"kafkaMq/database"
	kafka_tool "kafkaMq/kafkaTool"

	"github.com/segmentio/kafka-go"
)

type Customer struct {
	ctx    context.Context
	reader *kafka.Reader
	db     *sql.DB
}

func GetCustomer() *Customer {
	cfg := config.GetConfigure()
	return &Customer{
		ctx:    context.Background(),
		reader: kafka_tool.GetKafkaReadConn(cfg),
		db:     database.GetMysqlDBDriver(),
	}
}

func (c *Customer) Read() {
	defer func() {
		if err := c.reader.Close(); err != nil {
			panic(err)
		}
	}()
	for {
		m, err := c.reader.ReadMessage(c.ctx)
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset,
			string(m.Key), string(m.Value))
		// into database
		// sql := "insert into p_test(counts) values(1)"
		sql := string(m.Value)
		_, err = c.db.Exec(sql)
		if err != nil {
			panic(err)
		}
	}
}
