package main

import (
	"fmt"
	"kafkaMq/kafkaTool/customer"
	"kafkaMq/kafkaTool/producer"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// send()
	read()

}

func read() {
	reader := customer.GetCustomer()
	go reader.Read()

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
