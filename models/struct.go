package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Type of Migrate
type (
	Order struct {
		gorm.Model
		OrderId      string    `json:"order_id"`
		CustomerName string    `json:"customer_name"`
		OrderedAt    time.Time `json:"ordered_at"`
	}

	Item struct {
		gorm.Model
		ItemId      string `json:"item_id"`
		ItemCode    string `json:"item_code"`
		Description string `json:"description"`
		Quantity    int    `json:"quantity"`
		OrderId     string `json:"order_id"`
	}
)

// Type of Request
type (
	OrderRequest struct {
		CustomerName string        `json:"customerName"`
		OrderedAt    time.Time     `json:"orderedAt"`
		Items        []ItemRequest `json:"items"`
	}

	ItemRequest struct {
		LineItemId  string `json:"lineItemId"`
		ItemCode    string `json:"itemCode"`
		Description string `json:"description"`
		Quantity    int    `json:"quantity"`
	}
)

// Type of Data
type OrderData struct {
	Order
	Items []Item `json:"items"`
}
