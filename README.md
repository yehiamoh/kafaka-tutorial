# Kafka Producer-Consumer Demo with Sarama

## 📌 Overview

This project demonstrates a simple Kafka **Producer-Consumer** implementation using:

- **Sarama**: Go client library for Apache Kafka
- **Docker & Docker Compose**: For running Kafka and Zookeeper containers

---

## 📊 Kafka Architecture Diagram

                        +--------------------+
                         |     Producer       |
                         +---------+----------+
                                   |
                                   v
                     +-----------------------------+
                     |        Kafka Cluster        |
                     |  (Brokers / Topics / Parts) |
                     +-------------+---------------+
                                   |
                                   v
                         +--------------------+
                         |     Consumers      |
                         +--------------------+

                                 ↑
                                 |
              +----------------------------------------+
              |              Zookeeper                |
              |      (Cluster Coordination Service)   |
              +----------------------------------------+

yaml
Copy
Edit

---

## 🔑 Kafka Basics

- **Producer**: Publishes messages to Kafka topics
- **Consumer**: Reads messages from Kafka topics
- **Broker**: Kafka server that stores data
- **Topic**: Category/feed name to which messages are published
- **Partition**: Ordered, immutable sequence of messages within a topic
- **Zookeeper**: Coordinates Kafka brokers

---

## ✅ Prerequisites

- Docker & Docker Compose
- Go 1.24.1 or higher

---

## ⚙️ Setup

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/yehimoh/kafka-demo.git
cd kafka-demo
2. Start Kafka and Zookeeper
bash
Copy
Edit
docker-compose up -d
3. Create the Kafka Topic
bash
Copy
Edit
docker exec -it kafka /bin/bash

kafka-topics.sh --create \
  --topic=fancy-topic \
  --partitions=1 \
  --replication-factor=1 \
  --zookeeper=zookeeper:2181
📁 Project Structure
text
Copy
Edit
kafka-demo/
├── producer/
│   ├── main.go        # Kafka producer implementation
│   └── go.mod         # Go dependencies
├── consumer/
│   ├── main.go        # Kafka consumer implementation
│   └── go.mod         # Go dependencies
└── docker-compose.yml # Kafka and Zookeeper containers
💻 Key Code Components
▶️ Producer (producer/main.go)
go
Copy
Edit
config := sarama.NewConfig()
config.Producer.Return.Successes = true
config.Metadata.AllowAutoTopicCreation = false

conn, err := sarama.NewSyncProducer([]string{kafkaHost}, config)

msg := &sarama.ProducerMessage{
    Topic: topic,
    Key:   sarama.StringEncoder(fmt.Sprint(i)),
    Value: sarama.StringEncoder(fmt.Sprintf("Message :%v", i)),
}

partition, offset, err := conn.SendMessage(msg)
🔁 Consumer (consumer/main.go)
go
Copy
Edit
config := sarama.NewConfig()
config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
config.Consumer.Offsets.Initial = sarama.OffsetNewest

consumerGroup, err := sarama.NewConsumerGroup([]string{kafkaHost}, groupID, config)

func (h ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
    for msg := range claim.Messages() {
        fmt.Printf("Consumed message: %s\n", string(msg.Value))
        session.MarkMessage(msg, "")
    }
    return nil
}
🏃‍♂️ Running the Demo
1. Start the Consumer
bash
Copy
Edit
cd consumer
go run main.go
2. Start the Producer
bash
Copy
Edit
cd producer
go run main.go
📦 Dependencies
Sarama: Pure Go Kafka client library

Docker Images:

wurstmeister/zookeeper

wurstmeister/kafka

📈 Monitoring
Check topic details:

bash
Copy
Edit
docker exec -it kafka /bin/bash
kafka-topics.sh --describe --topic=fancy-topic --zookeeper=zookeeper:2181
🧹 Cleanup
Stop and remove containers:

bash
Copy
Edit
docker-compose down
❗ Troubleshooting
Connection Issues?

Verify Kafka is running:

bash
Copy
Edit
docker ps
Check Kafka logs:

bash
Copy
Edit
docker logs kafka
Topic Creation Fails?

Ensure Zookeeper is running first

Double-check the topic name matches in both Producer and Consumer
.
```
