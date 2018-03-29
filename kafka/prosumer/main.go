package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "192.168.240.86",
		"group.id":          "cclin",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	err = c.Subscribe("foo", nil)

	if err != nil {
		log.Fatal(err)
	}

	//

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "192.168.240.86"})
	if err != nil {
		log.Fatal(err)
	}

	defer p.Close()

	// Wait for message deliveries
	defer p.Flush(15 * 1000)

	//

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))

			bar := "bar"
			err := p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &bar, Partition: kafka.PartitionAny},
				Value:          msg.Value,
			}, nil)

			if err != nil {
				log.Fatalf("Producer error: %v (%v)\n", err, msg)
			}

			// break // just prosume one
		} else {
			log.Fatalf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
