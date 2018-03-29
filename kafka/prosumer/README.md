# Setup

```
brew install pkg-config
brew install librdkafka
go get -u github.com/confluentinc/confluent-kafka-go/kafka
```

# Usage

```
# terminal 1 - monitor topic bar
./kafka-console-consumer.sh --bootstrap-server 192.168.240.86:9092 --topic bar --from-beginning

# terminal 2 - consume message from topic foo then produce message to topic bar
go run main.go

# terminal 3 - produce message to topic foo
pip install kafka-python
python
>>> from kafka import KafkaProducer
>>> producer = KafkaProducer(bootstrap_servers='192.168.240.86:9092')
>>> producer.send('foo', b'1')

# terminal 2 - stop

# terminal 3 - produce message to topic foo

# terminal 2 - start
```

# Reference

* confluent-kafka-go https://github.com/confluentinc/confluent-kafka-go/
* confluent-kafka-go GoDoc https://godoc.org/github.com/confluentinc/confluent-kafka-go/kafka
