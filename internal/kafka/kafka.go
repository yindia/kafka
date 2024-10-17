package kafka

import (
	"fmt"
	v1 "kafka/internal/gen/cloud/v1"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde/protobuf"
)

type Producer struct {
	producer   *kafka.Producer
	serializer *protobuf.Serializer
}

// NewProducer initializes a new Kafka producer
func NewProducer(broker string) (*Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		return nil, err
	}
	client, err := schemaregistry.NewClient(schemaregistry.NewConfig("url"))

	if err != nil {
		fmt.Printf("Failed to create schema registry client: %s\n", err)
		return nil, err
	}

	ser, err := protobuf.NewSerializer(client, serde.ValueSerde, protobuf.NewSerializerConfig())

	if err != nil {
		fmt.Printf("Failed to create serializer: %s\n", err)
		return nil, err
	}

	return &Producer{producer: p, serializer: ser}, nil
}

// Send sends a message to the specified Kafka topic
func (p *Producer) Send(topic string, message interface{}) error {
	var payload []byte // Initialize payload

	switch msg := message.(type) {
	case v1.Profile:
		var err error
		payload, err = p.serializer.Serialize(topic, &msg)
		if err != nil {
			fmt.Printf("Failed to serialize payload: %s\n", err)
			return err
		}
	case v1.Repository:
		var err error
		payload, err = p.serializer.Serialize(topic, &msg)
		if err != nil {
			fmt.Printf("Failed to serialize payload: %s\n", err)
			return err
		}
	case v1.Commit:
		var err error
		payload, err = p.serializer.Serialize(topic, &msg)
		if err != nil {
			fmt.Printf("Failed to serialize payload: %s\n", err)
			return err
		}
	default:
		return fmt.Errorf("unsupported message type: %T", message) // Handle unsupported message types
	}

	// Produce the message
	deliveryChan := make(chan kafka.Event)
	err := p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: -1}, // Use -1 for any partition
		Value:          payload,
		Headers:        []kafka.Header{{Key: "test", Value: []byte("header values are binary")}},
	}, deliveryChan)

	if err != nil {
		return err
	}

	// Wait for delivery report
	go func() {
		e := <-deliveryChan
		m := e.(*kafka.Message)
		if m.TopicPartition.Error != nil {
			log.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		} else {
			log.Printf("Delivered message to %v [%d] at offset %v\n", *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}
	}()

	return nil
}

// Close closes the producer
func (p *Producer) Close() {
	p.producer.Flush(15 * 1000) // Wait for 15 seconds for messages to be delivered
	p.producer.Close()
}

type Consumer struct {
	consumer     *kafka.Consumer
	deserializer *protobuf.Deserializer
}

// NewConsumer initializes a new Kafka consumer
func NewConsumer(broker string, groupID string) (*Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          groupID,
		"auto.offset.reset": "earliest", // Start reading at the earliest message
	})
	if err != nil {
		return nil, err
	}

	client, err := schemaregistry.NewClient(schemaregistry.NewConfig("url"))

	if err != nil {
		fmt.Printf("Failed to create schema registry client: %s\n", err)
		return nil, err
	}

	deser, err := protobuf.NewDeserializer(client, serde.ValueSerde, protobuf.NewDeserializerConfig())

	if err != nil {
		fmt.Printf("Failed to create deserializer: %s\n", err)
		return nil, err
	}

	return &Consumer{consumer: c, deserializer: deser}, nil
}

func (c *Consumer) RegisterProfileMessage(message interface{}) error {
	return c.deserializer.ProtoRegistry.RegisterMessage((&v1.Profile{}).ProtoReflect().Type())
}

func (c *Consumer) RegisterRepositoryMessage(message interface{}) error {
	return c.deserializer.ProtoRegistry.RegisterMessage((&v1.Repository{}).ProtoReflect().Type())
}

func (c *Consumer) RegisterCommitMessage(message interface{}) error {
	return c.deserializer.ProtoRegistry.RegisterMessage((&v1.Commit{}).ProtoReflect().Type())
}

// Consume consumes messages from the specified Kafka topic
func (c *Consumer) Consume(topic string) error {
	err := c.consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return err
	}

	for {
		msg, err := c.consumer.ReadMessage(-1) // -1 means wait indefinitely for a message
		if err == nil {
			log.Printf("Received message: %s from topic: %s\n", string(msg.Value), *msg.TopicPartition.Topic)
		} else {
			log.Printf("Error while consuming message: %v\n", err)
		}
	}
}

// Close closes the consumer
func (c *Consumer) Close() {
	c.consumer.Close()
}
