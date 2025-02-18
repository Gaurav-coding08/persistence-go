package services

import (
	"encoding/json"
	"log"

	repoModels "github.com/Gaurav-coding08/persistence-go/internal/app/repositories/models"
)

type Repo interface {
	UpdateStock(transaction *repoModels.StockDetail) error
}

type Service struct {
	repo Repo
}

func New(repo Repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) HandleEvent(payload json.RawMessage) error {
	return s.handleStockPriceUpdate(payload)
}

// handleStockPriceUpdate processes stock price update events
func (s *Service) handleStockPriceUpdate(payload json.RawMessage) error {
	// in place of internally declare, use global package
	var stockReq StockDetail
	if err := json.Unmarshal(payload, &stockReq); err != nil {
		log.Printf("Failed to parse stock update: %v", err)
		return err
	}

	stock := &repoModels.StockDetail{
		ID:    stockReq.ID,
		Name:  stockReq.Name,
		Price: stockReq.Price,
	}

	err := s.repo.UpdateStock(stock)
	if err != nil {
		log.Printf("Error updating stock price: %v", err)
		return err
	}

	return nil
}
