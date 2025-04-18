package main

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

const (
	kafkaHost = "localhost:9092"
	topic     = "fancy-topic"
)

func main() {
	config := sarama.NewConfig() // Client Libraray for Apache Kafka
	config.Producer.Return.Successes=true
	config.Metadata.AllowAutoTopicCreation=false

	conn,err:=sarama.NewSyncProducer([]string{kafkaHost},config)
	if err!=nil{
		log.Fatal("Couldn't connect to Kafaka",err)
	}
	defer conn.Close()

	i:=0 
	for{
		msg:=&sarama.ProducerMessage{
			Topic: topic, //Topic For Kafka
			Key: sarama.StringEncoder(fmt.Sprint(i)), // ID for the message used in partitioning
			Value: sarama.StringEncoder(fmt.Sprintf("Message :%v",i)), // The value of the message
		}
		partition,offset,err:=conn.SendMessage(msg) // sending message to the Topic
		if err!=nil{
			log.Fatal("Couldn't send message",err)
		}
		log.Printf("Sent Message :%v,partition:%v,offset:%v",msg.Value,partition,offset)
		i++
		time.Sleep(time.Second)
	}
}/*
 docker exec -it kafka /bin/bash
 kafka-topics.sh --create --topic=fancy-topic --partitions=1 --replication-factor=1 --zookeeper=zookeeper:2181
 creating Kafka Topic

kafka-topics.sh --describe --topic=fancy-topic --zookeeper=zookeeper:2181
describe the topic
 */