package main

import (
    "fmt"
    "log"
    "github.com/confluentinc/confluent-kafka-go/kafka"
)

func consumer() {

    c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost",
        "group.id":          "myGroup",
        "auto.offset.reset": "earliest",
    })

    if err != nil {
        panic(err)
    }

    c.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)

    for {
        msg, err := c.ReadMessage(-1)
        if err == nil {
            fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
        } else {
            fmt.Printf("Consumer error: %v (%v)\n", err, msg)
            break
        }
    }

    c.Close()
}

func producer() {

    p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
    if err != nil {
        panic(err)
    }

    // Delivery report handler for produced messages
    go func() {
        for e := range p.Events() {
            switch ev := e.(type) {
            case *kafka.Message:
                if ev.TopicPartition.Error != nil {
                    fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
                } else {
                    fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
                }
            }
        }
    }()

    // Produce messages to topic (asynchronously)
    topic := "myTopic"
    for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
        p.Produce(&kafka.Message{
            TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
            Value:          []byte(word),
        }, nil)
    }

    // Wait for message deliveries
    p.Flush(15 * 1000)
}

func main() {
    log.Println("TODO")
}
