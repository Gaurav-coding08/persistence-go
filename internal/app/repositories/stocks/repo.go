package stocks

import (
	"log"

	models "github.com/Gaurav-coding08/persistence-go/internal/app/repositories/models"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) UpdateStock(updatedStock *models.StockDetail) error {
	var existingStock models.StockDetail

	result := r.db.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)}).
		Where("id = ?", updatedStock.ID).
		First(&existingStock)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			if err := r.db.Create(updatedStock).Error; err != nil {
				log.Printf("Failed to insert new stock: %v", err)
				return err
			}

		} else {
			return result.Error
		}
	} else {
		err := r.db.Model(&existingStock).Updates(models.StockDetail{
			Price:     updatedStock.Price,
			UpdatedAt: updatedStock.UpdatedAt,
		}).Error

		if err != nil {

			return err
		}

	}

	return nil
}
