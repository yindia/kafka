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

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)                  // Include date, time, and file info in logs
	log.Println("Logging initialized: All logs will be printed.") // Log initialization message
}

type Producer struct {
	producer   *kafka.Producer
	serializer *protobuf.Serializer
}

// NewProducer initializes a new Kafka producer with the specified broker and schema registry URL.
func NewProducer(broker string, schemaRegistryURL string) (*Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker, // Use the provided broker address
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}

	client, err := schemaregistry.NewClient(schemaregistry.NewConfig(schemaRegistryURL))
	if err != nil {
		return nil, fmt.Errorf("failed to create schema registry client: %w", err)
	}

	ser, err := protobuf.NewSerializer(client, serde.ValueSerde, protobuf.NewSerializerConfig())
	if err != nil {
		return nil, fmt.Errorf("failed to create serializer: %w", err)
	}

	return &Producer{producer: p, serializer: ser}, nil
}

// sendMessage is a helper function to send a serialized message to the specified Kafka topic.
func (p *Producer) sendMessage(topic string, payload []byte) error {
	deliveryChan := make(chan kafka.Event)
	err := p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic}, // Use -1 for any partition
		Value:          payload,
		Headers:        []kafka.Header{{Key: "test", Value: []byte("header values are binary")}},
	}, deliveryChan)

	if err != nil {
		return fmt.Errorf("failed to produce message: %w", err)
	}

	// Wait for delivery report
	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		return fmt.Errorf("delivery failed: %v", m.TopicPartition.Error)
	}
	log.Printf("Delivered message to %v [%d] at offset %v\n", *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)

	p.producer.Flush(15 * 1000) // Wait for 15 seconds for messages to be delivered
	return nil
}

// SendProfile sends a serialized Protobuf Profile message to the specified Kafka topic.
func (p *Producer) SendProfile(topic string, message *v1.Profile) error {
	payload, err := p.serializer.Serialize(topic, message) // Serialize the Protobuf message
	if err != nil {
		return fmt.Errorf("failed to serialize payload: %w", err)
	}
	return p.sendMessage(topic, payload) // Use the helper function
}

// SendRepository sends a serialized Protobuf Repository message to the specified Kafka topic.
func (p *Producer) SendRepository(topic string, message *v1.Repository) error {
	payload, err := p.serializer.Serialize(topic, message) // Serialize the Protobuf message
	if err != nil {
		return fmt.Errorf("failed to serialize payload: %w", err)
	}
	return p.sendMessage(topic, payload) // Use the helper function
}

// Close closes the producer and flushes any remaining messages.
func (p *Producer) Close() {
	p.producer.Flush(15 * 1000) // Wait for 15 seconds for messages to be delivered
	p.producer.Close()
}

type Consumer struct {
	consumer     *kafka.Consumer
	deserializer *protobuf.Deserializer
}

// NewConsumer initializes a new Kafka consumer with the specified broker and group ID.
func NewConsumer(broker string, groupID string) (*Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          groupID,
		"auto.offset.reset": "earliest", // Start reading at the earliest message
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer: %w", err)
	}

	client, err := schemaregistry.NewClient(schemaregistry.NewConfig("url"))
	if err != nil {
		return nil, fmt.Errorf("failed to create schema registry client: %w", err)
	}

	deser, err := protobuf.NewDeserializer(client, serde.ValueSerde, protobuf.NewDeserializerConfig())
	if err != nil {
		return nil, fmt.Errorf("failed to create deserializer: %w", err)
	}

	return &Consumer{consumer: c, deserializer: deser}, nil
}

func (c *Consumer) RegisterProfileMessage(message interface{}) error {
	return c.deserializer.ProtoRegistry.RegisterMessage((&v1.Profile{}).ProtoReflect().Type())
}

func (c *Consumer) RegisterRepositoryMessage(message interface{}) error {
	return c.deserializer.ProtoRegistry.RegisterMessage((&v1.Repository{}).ProtoReflect().Type())
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
