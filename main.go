package main

import (
	"fmt"
	"kafkaMq/config"
	kafka_customer "kafkaMq/kafkaTool/customer"
	kafka_producer "kafkaMq/kafkaTool/producer"
	mqtt_tool "kafkaMq/mqttTool"
	mqtt_producer "kafkaMq/mqttTool/producer"
	"log"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/segmentio/kafka-go"
)

func main() {
	/*
		性能测试整体思路-暴力测试
		1. 生产出信息发送至mqtt（一万条）
		2. mqtt消费后，发送至kafka
		3. kafka收到消费，读取并入库
		4. 读取入库结果
	*/
	cfg := config.GetConfigure()
	clientID := "performance" // mqtt client_id
	mqtt_topic := "crazy1"
	kafka_topic := "crazy"
	// mqtt client
	client := mqtt_tool.GetClient(cfg, clientID)
	// 消费数据发送到kafka
	go mqttCustomer(client, mqtt_topic, kafka_topic)
	// 生产数据
	// go mqttProduce(client, mqtt_topic)
	// 消费kafka数据到数据库
	go kafkaConsume()

	// 异常通知
	s := waitForSignal()
	log.Println("Signal received, server close, ", s)
}

// generate message to mqtt
func mqttProduce(client mqtt.Client, topic string) {
	// 完成后 2s关闭mqtt client
	// defer func(client mqtt.Client) {
	// 	// 完成后 2s关闭mqtt client
	// 	client.Disconnect(2000)
	// 	log.Println("mqtt client closed...,exiting...")
	// }(client)
	i := 0
	total := 50
	for i = 0; i <= total; i++ {
		// go func(client mqtt.Client, i int, topic string) {
		// 	if ack := mqtt_producer.PublishMessage(client, fmt.Sprintf("insert into p_test(counts) values(%d)", i), topic,
		// 		2, false); !ack {
		// 		log.Printf("mqtt publish[%d] is failed....", i)
		// 		// continue
		// 	}
		// 	log.Printf("mqtt publish[%d] is succeed....", i)
		// }(client, ii, topic)

		if ack := mqtt_producer.PublishMessage(client, fmt.Sprintf("insert into p_test(counts) values(%d)", i), topic,
			1, false); !ack {
			log.Printf("mqtt publish[%d] is failed....", i)
			// continue
		}
		log.Printf("mqtt publish[%d] is succeed....", i)
	}
}

// customer message from mqtt to kafka
func mqttCustomer(client mqtt.Client, mq_topic, kafka_topic string) {
	// timer := time.NewTicker(1 * time.Second)
	kafkaProducer := kafka_producer.GetProducer()
	if token := client.Subscribe(mq_topic, 1, func(c mqtt.Client, m mqtt.Message) {
		if err := kafkaProducer.Send([]kafka.Message{{
			Topic: kafka_topic,
			Value: m.Payload(),
		}}...); err != nil {
			log.Printf("kafka produce failed, err is %s", err.Error())
		}
	}); token.Wait() && token.Error() != nil {
		log.Printf("mqtt customer failed, err is %s", token.Error())
		kafkaProducer.Writer.Close()
		return
	}
	log.Printf("mqtt customer success....")
}

// consume kafka to database
func kafkaConsume() {
	consumer := kafka_customer.GetCustomer()
	consumer.Read()
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

// func send() {
func main1() {
	// topics := []string{"orderInfo", "test1", "test2"}
	// topics := []string{"mp"}
	sendCount := 10
	p := kafka_producer.GetProducer()
	defer p.Writer.Close()
	for i := 0; i < sendCount; i++ {
		message := make([]kafka.Message, 0)
		message = append(
			message, kafka.Message{
				// Topic: "orderInfo",
				Topic: "mp",
				//分区是只读的，写消息时绝对不能设置分区
				// Partition: 0,
				// Value: []byte(fmt.Sprintf("hello kafka[%d]", i)),
				Value: []byte("insert into p_test(counts) values(1)"),
				// Value: []byte(fmt.Sprintf("orderID[%s]:已发货", order[i])),
				// Key:   []byte("A"),
				// Key: []byte(order[i]),
			})
		// send
		if err := p.Send(message...); err != nil {
			panic(err)
		}
	}
	log.Print("send succeed")
}
