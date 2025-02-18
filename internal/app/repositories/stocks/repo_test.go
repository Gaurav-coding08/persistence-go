package stocks_test

import (
	"log"
	"testing"
	"time"

	models "github.com/Gaurav-coding08/persistence-go/internal/app/repositories/models"
	repo "github.com/Gaurav-coding08/persistence-go/internal/app/repositories/stocks"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.StockDetail{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestUpdateStock_InsertNewStock(t *testing.T) {
	db, err := setupTestDB()
	require.NoError(t, err)

	repo := repo.New(db)

	newStock := &models.StockDetail{
		ID:        1,
		Name:      "Test Stock",
		Price:     100.50,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = repo.UpdateStock(newStock)
	require.NoError(t, err)

	var storedStock models.StockDetail
	err = db.First(&storedStock, "id = ?", newStock.ID).Error
	require.NoError(t, err)
	require.Equal(t, newStock.Price, storedStock.Price)
}

func TestUpdateStock_UpdateExistingStock(t *testing.T) {
	db, err := setupTestDB()
	require.NoError(t, err)

	repo := repo.New(db)

	initialStock := &models.StockDetail{
		ID:        1,
		Name:      "Test Stock",
		Price:     100.00,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = db.Create(initialStock).Error
	require.NoError(t, err)

	updatedStock := &models.StockDetail{
		ID:        1,
		Name:      "Test Stock",
		Price:     150.75,
		UpdatedAt: time.Now(),
	}
	err = repo.UpdateStock(updatedStock)
	require.NoError(t, err)

	var storedStock models.StockDetail
	err = db.First(&storedStock, "id = ?", updatedStock.ID).Error
	require.NoError(t, err)
	require.Equal(t, updatedStock.Price, storedStock.Price)
}

func TestUpdateStock_DBError(t *testing.T) {
	db, err := setupTestDB()
	require.NoError(t, err)

	repo := repo.New(db)

	sqlDB, _ := db.DB()
	sqlDB.Close()

	newStock := &models.StockDetail{
		ID:        2,
		Name:      "Broken Stock",
		Price:     99.99,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = repo.UpdateStock(newStock)
	require.Error(t, err)
	log.Printf("Expected DB error: %v", err)
}
