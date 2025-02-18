package server

import (
	"log"

	consumerSvc "github.com/Gaurav-coding08/persistence-go/internal/app/services/consumers"
	eventsSvc "github.com/Gaurav-coding08/persistence-go/internal/app/services/events"

	stocksRepo "github.com/Gaurav-coding08/persistence-go/internal/app/repositories/stocks"
	stocksSvc "github.com/Gaurav-coding08/persistence-go/internal/app/services/stocks"

	"github.com/Gaurav-coding08/persistence-go/config"
	"gorm.io/gorm"
)

func StartConsumers(cfg *config.AppConfig, db *gorm.DB) {

	stocksSvc := stocksSvc.New(stocksRepo.New(db))

	eventService := eventsSvc.New(stocksSvc)

	eventConsumer := consumerSvc.New(cfg.KafkaConfig.Broker, cfg.KafkaConfig.Topic, eventService)

	// Start Consumer in a Goroutine
	go eventConsumer.Consume()

	log.Println("Kafka Consumers Started Successfully!")

}
