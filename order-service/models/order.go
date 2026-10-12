package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID     uint    `json:"user_id"`
	ProductID  uint    `json:"product_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
	Status     string  `json:"status"` // PENDING, COMPLETED, CANCELLED
}