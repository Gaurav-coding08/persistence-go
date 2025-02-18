package models

import "time"

type StockDetail struct {
	ID        uint       `gorm:"primaryKey"`
	Name      string     `gorm:"column:name"`
	Price     float64    `gorm:"column:price"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

func (StockDetail) TableName() string {
	return "stock_details"
}
