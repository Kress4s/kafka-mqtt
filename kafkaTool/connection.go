package kafka_tool

import (
	"kafkaMq/config"
	"strings"
	"sync"

	"github.com/segmentio/kafka-go"
)

var (
	conn kafka.Writer
	once sync.Once
)

var topics = []string{"orderInfo", "test1", "test2"}

func GetKafkaWriteConn(cfg *config.Configure) *kafka.Writer {
	once.Do(func() {
		conn = kafka.Writer{
			Addr: kafka.TCP(strings.Join([]string{cfg.Kafka.Addr, cfg.Kafka.Port}, ":")),
			// NOTE: When Topic is not defined here, each Message must define it instead.
			// Topic:    topic,

			// LeastBytes 至少分配一次，顺序分配
			Balancer: &kafka.LeastBytes{},

			// Hash 哈希分配，实现相同的key能分发到相同topic中同一个partition中
			// Balancer: &kafka.Hash{},
		}
	})
	// rewrite single conn for different topic by message provide
	return &conn
}

func GetKafkaReadConn(cfg *config.Configure) *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		GroupID:     "t1",
		GroupTopics: topics,
		Brokers:     []string{strings.Join([]string{cfg.Kafka.Addr, cfg.Kafka.Port}, ":")},
		// StartOffset: 0,
	})
	return r
}