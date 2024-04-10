package main

import (
	"fmt"
	"kafkaMq/config"
	"kafkaMq/kafkaTool/producer"
	mqtt_producer "kafkaMq/mqttTool/producer"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	cfg := config.GetConfigure()
	mqttPublisher := mqtt_producer.GetPublisher(cfg)
	if token := mqttPublisher.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
		os.Exit(1)
	}
	defer mqttPublisher.Disconnect(250)

	// 订阅主题
	if token := mqttPublisher.Subscribe("order", 2, nil); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
		os.Exit(1)
	}
	// count := 2
	// for i := 0; i < count; i++ {
	// 	token := mqttPublisher.Publish("order", 1, false, fmt.Sprintf("%d member", i))
	// 	token.Wait()
	// }

	// 处理系统信号，以便在接收到SIGINT或SIGTERM时优雅地关闭程序
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	<-signalChan
	fmt.Println("Received signal, shutting down...")
}

func read() {
	// reader := customer.GetCustomer()

	// go reader.Read()

	// go send()

	time.Sleep(100 * time.Second)

}

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

// func waitForSignal() os.Signal {
// 	signalChan := make(chan os.Signal, 1)
// 	defer close(signalChan)
// 	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
// 	s := <-signalChan
// 	signal.Stop(signalChan)
// 	return s
// }
