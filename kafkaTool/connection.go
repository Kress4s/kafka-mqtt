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

// var topics = []string{"orderInfo", "test1", "test2"}
var topics = []string{"crazy"}

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
			RequiredAcks: kafka.RequireOne,
		}
	})
	// rewrite single conn for different topic by message provide
	return &conn
}

func GetKafkaReadConn(cfg *config.Configure) *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		GroupID:     "cc1",
		GroupTopics: topics,
		// 多个broker
		Brokers: []string{strings.Join([]string{cfg.Kafka.Addr, cfg.Kafka.Port}, ":"), "121.41.38.13:29092",
			"121.41.38.13:39092"},
		// 当一个特定的partition没有commited offset时(比如第一次读一个partition，之前没有commit过)，
		// 通过StartOffset指定从第一个还是最后一个位置开始消费。StartOffset的取值要么是FirstOffset要么是LastOffset，
		// LastOffset表示Consumer启动之前生成的老数据不管了。
		// 仅当指定了GroupID时，StartOffset才生效。
		StartOffset: kafka.FirstOffset,
	})
	return r
}
