package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

const (
	kafakaHost = "localhost:9092"
	topic      = "fancy-topic"
	groupID    = "consumer-group1"
)

type JsonMessage struct{
	ID int `json:"id"`
	Content string `json:"content"`

}

func main() {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy=sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial=sarama.OffsetNewest
	config.Version = sarama.V2_1_0_0

	consumerGroup,err:=sarama.NewConsumerGroup([]string{kafakaHost},groupID,config)
	if err!=nil{
		log.Fatal("Error creating consumer group: ", err)
	}
	defer consumerGroup.Close()

	handler:= ConsumerGroupHandler{}
	fmt.Printf("Listening to topic %s...\n", topic)
	for {
		err := consumerGroup.Consume(context.TODO(), []string{topic}, handler)
		if err != nil {
			log.Println("Error consuming: ", err)
			time.Sleep(time.Second)
		}
	}

}
// ConsumerGroupHandler implements sarama.ConsumerGroupHandler
type ConsumerGroupHandler struct{}

func (h ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (h ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		val:=&JsonMessage{}
		err:=json.Unmarshal(msg.Value,val)
		if err!=nil{
			log.Fatal("Couldn't decode the json :",err)
		}
		fmt.Printf("Consumed message: %v, partition: %d, time: %s\n",
			val, msg.Partition, time.Now().Format("05:000"))
		session.MarkMessage(msg, "")
		time.Sleep(time.Second)
	}
	return nil
}