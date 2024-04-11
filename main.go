package main

import (
	"fmt"
	"kafkaMq/config"
	"kafkaMq/kafkaTool/producer"
	mqtt_tool "kafkaMq/mqttTool"
	mqtt_customer "kafkaMq/mqttTool/customer"
	mqtt_producer "kafkaMq/mqttTool/producer"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/segmentio/kafka-go"
)

func main() {
	cfg := config.GetConfigure()
	clientID := "123"
	client := mqtt_tool.GetClient(cfg, clientID)
	defer mqtt_tool.DisClientConnect(client, 250)
	mqtt_customer.Subscriber(client, "order", 2)
	count := 2
	for i := 0; i < count; i++ {
		mqtt_producer.PublishMessage(client, fmt.Sprintf("%d member", i), "order", 1, false)
	}
	s := waitForSignal()
	fmt.Printf("Received signal,[%+v], shutting down...", s)
}

func waitForSignal() os.Signal {
	// 处理系统信号，以便在接收到SIGINT或SIGTERM时优雅地关闭程序
	signalChan := make(chan os.Signal, 1)
	defer close(signalChan)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	s := <-signalChan
	signal.Stop(signalChan)
	return s
}

// func read() {
// reader := customer.GetCustomer()

// go reader.Read()

// go send()

// time.Sleep(100 * time.Second)

// }

func send() {
	topics := []string{"orderInfo", "test1", "test2"}
	sendCount := 100
	message := make([]kafka.Message, 0, sendCount)
	for i := 0; i < sendCount; i++ {
		message = append(
			message, kafka.Message{
				// Topic: "orderInfo",
				Topic: topics[i%3],
				//分区是只读的，写消息时绝对不能设置分区
				// Partition: 0,
				Value: []byte(fmt.Sprintf("hello kafka[%d]", i)),
				// Value: []byte(fmt.Sprintf("orderID[%s]:已发货", order[i])),
				// Key:   []byte("A"),
				// Key: []byte(order[i]),
			})
	}

	// send
	if err := producer.GetProducer().Send(message...); err != nil {
		panic(err)
	}
	log.Print("send succeed")
}
