package oapi_test

import (
	"testing"
	"time"

	"receipt-processor-challenge/internal/delivery/oapi"
	"receipt-processor-challenge/internal/domain"

	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/stretchr/testify/assert"
)

func TestReceiptAndItemToDTO(t *testing.T) {
	baseDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local)
	expectedTime := time.Date(2022, 1, 1, 13, 1, 0, 0, time.Local)

	tests := []struct {
		name          string
		receipt       oapi.Receipt
		expectedDTO   domain.ReceiptDTO
		expectedError bool
	}{
		{
			name: "Valid Receipt",
			receipt: oapi.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: openapi_types.Date{Time: baseDate},
				PurchaseTime: "13:01",
				Items: []oapi.Item{
					{
						ShortDescription: "Item 1",
						Price:            "1.00",
					},
					{
						ShortDescription: "Item 2",
						Price:            "1.00",
					},
				},
				Total: "2.00",
			},
			expectedDTO: domain.ReceiptDTO{
				Retailer:    "M&M Corner Market",
				PurchasedAt: expectedTime,
				Items: []domain.ItemDTO{
					{
						ShortDescription: "Item 1",
						Price:            1.00,
					},
					{
						ShortDescription: "Item 2",
						Price:            1.00,
					},
				},
				Total: 2.00,
			},
		},
		{
			name: "Invalid Receipt Total",
			receipt: oapi.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: openapi_types.Date{Time: baseDate},
				PurchaseTime: "13:01",
				Items: []oapi.Item{
					{
						ShortDescription: "Item 1",
						Price:            "1.00",
					},
				},
				Total: "Invalid Total",
			},
			expectedError: true,
		},
		{
			name: "Invalid Receipt Time",
			receipt: oapi.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: openapi_types.Date{Time: baseDate},
				PurchaseTime: "Invalid Time",
				Items: []oapi.Item{
					{
						ShortDescription: "Item 1",
						Price:            "1.00",
					},
				},
				Total: "1.00",
			},
			expectedError: true,
		},
		// Additionally tests oapi.Item.ToDTO()
		{
			name: "Invalid Item Price",
			receipt: oapi.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: openapi_types.Date{Time: baseDate},
				PurchaseTime: "Invalid Time",
				Items: []oapi.Item{
					{
						ShortDescription: "Item 1",
						Price:            "Invalid Price",
					},
				},
				Total: "1.00",
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dto, err := tt.receipt.ToDTO()
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedDTO, dto)
			}
		})
	}
}
