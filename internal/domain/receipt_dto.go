package domain

import "time"

// Receipt defines model for Receipt.
type ReceiptDTO struct {
	Items       []ItemDTO
	PurchasedAt time.Time
	Retailer    string
	Total       float64
}
