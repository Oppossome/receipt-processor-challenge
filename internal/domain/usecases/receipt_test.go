package usecases_test

import (
	"context"
	"testing"
	"time"

	"receipt-processor-challenge/internal/domain"
	"receipt-processor-challenge/internal/domain/usecases"

	"github.com/stretchr/testify/assert"
)

// If you're curious what the purpose of `MARK:` is see: https://code.visualstudio.com/docs/getstarted/userinterface#_minimap
// MARK: TestProcessReceipt
func TestProcessReceipt(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// An item not eligble for any points.
	pointlessItem := domain.ItemDTO{
		ShortDescription: "abcd", // Length is not a multiple of 3
		Price:            1.23,   // Not a multiple of 0.25
	}

	// A receipt not eligble for any points.
	pointlessReceipt := domain.ReceiptDTO{
		Retailer:    "!!!",                                         // No alphanumeric characters
		PurchasedAt: time.Date(2022, 3, 20, 13, 0, 0, 0, time.UTC), // Day is even, time is not between 2:00pm and 4:00pm
		Items:       []domain.ItemDTO{pointlessItem},
		Total:       1.23, // Not a round dollar amount
	}

	tests := []struct {
		name    string
		points  int64
		receipt domain.ReceiptDTO
	}{
		// MARK: - PointlessReceipt
		// !! If this test is failing, it has ramifications for every other test.
		{
			name:   "Pointless Receipt",
			points: 0,
			// Reconstruct it in a fashion typical to the tests so I can simply copy it for editing into other tests 😄
			receipt: domain.ReceiptDTO{
				Retailer:    pointlessReceipt.Retailer,
				PurchasedAt: pointlessReceipt.PurchasedAt,
				Items:       pointlessReceipt.Items,
				Total:       pointlessReceipt.Total,
			},
		},

		// MARK: - Rule 1.
		// 1. One point for every alphanumeric character in the retailer name.
		{
			name:   "Alphanumeric character points",
			points: 4,
			receipt: domain.ReceiptDTO{
				// Include _ to ensure we're not inadvertently using \w instead of [A-Za-z0-9]
				Retailer:    "a_b-c d",
				PurchasedAt: pointlessReceipt.PurchasedAt,
				Items:       pointlessReceipt.Items,
				Total:       pointlessReceipt.Total,
			},
		},

		// MARK: - Rule 2 & 3.
		// 2. 50 points if the total is a round dollar amount with no cents
		// 3. 25 points if the total is a multiple of `0.25`
		{
			name:   "Multiple of 1.00 and 0.25 points",
			points: 75,
			receipt: domain.ReceiptDTO{
				Retailer:    pointlessReceipt.Retailer,
				PurchasedAt: pointlessReceipt.PurchasedAt,
				Items:       pointlessReceipt.Items,
				Total:       1.00,
			},
		},
		{
			name:   "Multiple of 0.25 points",
			points: 25,
			receipt: domain.ReceiptDTO{
				Retailer:    pointlessReceipt.Retailer,
				PurchasedAt: pointlessReceipt.PurchasedAt,
				Items:       pointlessReceipt.Items,
				Total:       1.25,
			},
		},

		// MARK: - Rule 4.
		// 4. 5 points for every two items on the receipt.
		{
			name:   "5 points for every two items points",
			points: 10,
			receipt: domain.ReceiptDTO{
				Retailer:    pointlessReceipt.Retailer,
				PurchasedAt: pointlessReceipt.PurchasedAt,
				Items: []domain.ItemDTO{
					pointlessItem,
					pointlessItem,
					pointlessItem,
					pointlessItem,
				},
				Total: pointlessItem.Price,
			},
		},

		// MARK: - Rule 5.
		// Rule 5. If the trimmed length of the item description is a multiple of 3, multiply the price by `0.2` and round up to the nearest integer. The result is the number of points earned.
		{
			name:   "Item description is a multiple of 3 points",
			points: 10,
			receipt: domain.ReceiptDTO{
				Retailer:    pointlessReceipt.Retailer,
				PurchasedAt: pointlessReceipt.PurchasedAt,
				Items: []domain.ItemDTO{
					{
						ShortDescription: "abc",
						// Using 46 because it times 0.2 is closer to 9 than 10, ensuring we aren't simply math.Round-ing
						// 46 * 0.2 = 9.2 (Which rounds up to 10)
						Price: 46,
					},
				},
				Total: pointlessItem.Price,
			},
		},

		// MARK: - Rule 6.
		// Rule 6. 6 points if the day in the purchase date is odd.
		{
			name:   "Odd day purchase points",
			points: 6,
			receipt: domain.ReceiptDTO{
				Retailer:    pointlessReceipt.Retailer,
				PurchasedAt: time.Date(2022, 3, 21, 0, 0, 0, 0, time.UTC),
				Items:       pointlessReceipt.Items,
				Total:       pointlessReceipt.Total,
			},
		},

		// MARK: - Rule 7.
		// Rule 7. 10 points if the time of purchase is after 2:00pm and before 4:00pm
		{
			name:   "Purchase between 2PM and 4PM points",
			points: 10,
			receipt: domain.ReceiptDTO{
				Retailer:    pointlessReceipt.Retailer,
				PurchasedAt: time.Date(2022, 3, 20, 14, 1, 0, 0, time.Local),
				Items:       pointlessReceipt.Items,
				Total:       pointlessReceipt.Total,
			},
		},
		{
			name:   "No points for a purchase exactly at 2PM",
			points: 0,
			receipt: domain.ReceiptDTO{
				Retailer:    pointlessReceipt.Retailer,
				PurchasedAt: time.Date(2022, 3, 20, 14, 0, 0, 0, time.Local),
				Items:       pointlessReceipt.Items,
				Total:       pointlessReceipt.Total,
			},
		},
		{
			name:   "No points for a purchase exactly at 4PM",
			points: 0,
			receipt: domain.ReceiptDTO{
				Retailer:    pointlessReceipt.Retailer,
				PurchasedAt: time.Date(2022, 3, 20, 16, 0, 0, 0, time.Local),
				Items:       pointlessReceipt.Items,
				Total:       pointlessReceipt.Total,
			},
		},

		// MARK: - Fetch Examples
		// These were provided by Fetch and are useful as a concrete truth
		{
			name:   "Provided Example 1",
			points: 109,
			receipt: domain.ReceiptDTO{
				Retailer:    "M&M Corner Market",
				PurchasedAt: time.Date(2022, 3, 20, 14, 33, 0, 0, time.Local),
				Items: []domain.ItemDTO{
					{
						ShortDescription: "Gatorade",
						Price:            2.25,
					},
					{
						ShortDescription: "Gatorade",
						Price:            2.25,
					},
					{
						ShortDescription: "Gatorade",
						Price:            2.25,
					},
					{
						ShortDescription: "Gatorade",
						Price:            2.25,
					},
				},
				Total: 9.00,
			},
		},
		{
			name:   "Provided Example 2",
			points: 28,
			receipt: domain.ReceiptDTO{
				Retailer:    "Target",
				PurchasedAt: time.Date(2022, 1, 1, 13, 1, 0, 0, time.Local),
				Items: []domain.ItemDTO{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            6.49,
					},
					{
						ShortDescription: "Emils Cheese Pizza",
						Price:            12.25,
					},
					{
						// I unironically eat these like a couple of times a week 😂
						ShortDescription: "Knorr Creamy Chicken",
						Price:            1.26,
					},
					{
						ShortDescription: "Doritos Nacho Cheese",
						Price:            3.35,
					},
					{
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						Price:            12.00,
					},
				},
				Total: 35.35,
			},
		},
	}

	// MARK: - Test Runner
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecases := usecases.NewUsecases()

			// Process our receipt, which calculates and stores the points associated with it
			id := usecases.ProcessReceipt(ctx, tt.receipt)

			// Retrieve the points and check 'em
			points, err := usecases.GetReceiptPoints(ctx, id)

			assert.NoError(t, err)
			assert.Equal(t, tt.points, points)
		})
	}
}
