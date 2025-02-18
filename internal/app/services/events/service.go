package services

import (
	"encoding/json"
	"log"

	stockSvc "github.com/Gaurav-coding08/persistence-go/internal/app/services/stocks"
	"github.com/segmentio/kafka-go"
)

type EventService struct {
	StocksService *stockSvc.Service
}

func New(stocksSvc *stockSvc.Service) *EventService {
	return &EventService{
		StocksService: stocksSvc,
	}
}

func (e *EventService) ProcessMessage(rawMessage kafka.Message) {
	var kafkaMessage KafkaMessage

	if err := json.Unmarshal(rawMessage.Value, &kafkaMessage); err != nil {
		log.Printf(" Failed to parse Kafka message: %v", err)
		return
	}

	log.Printf("Processing Event: %s", kafkaMessage.EventType)

	switch kafkaMessage.EventType {
	case StockUpdateEvent:
		e.StocksService.HandleEvent(kafkaMessage.Payload)

	default:
		log.Printf("Unknown event type: %s", kafkaMessage.EventType)
	}
}
