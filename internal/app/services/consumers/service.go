package services

import (
	"context"
	"log"

	eventsSvc "github.com/Gaurav-coding08/persistence-go/internal/app/services/events"
	"github.com/segmentio/kafka-go"
)

type EventConsumer struct {
	reader       *kafka.Reader
	eventService *eventsSvc.EventService
}

func New(
	broker string,
	topic string,
	eventService *eventsSvc.EventService) *EventConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: "consumer-1-group",
	})

	return &EventConsumer{
		reader,
		eventService}
}

func (c *EventConsumer) Consume() {

	for {
		msg, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("❌ Error reading Kafka message: %v", err)
			continue
		}

		c.eventService.ProcessMessage(msg)

		// Commit only if processing was successful
		err = c.reader.CommitMessages(context.Background(), msg)
		if err != nil {
			log.Printf("⚠️ Failed to commit Kafka offset: %v", err)
		}
	}
}
