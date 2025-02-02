package usecases

import (
	"context"
	"fmt"
	"math"
	"regexp"
	"strings"

	"receipt-processor-challenge/internal/domain"

	"github.com/google/uuid"
)

var alphanumericRegexp = regexp.MustCompile("[A-Za-z0-9]")

func (u *Usecases) ProcessReceipt(ctx context.Context, receipt domain.ReceiptDTO) uuid.UUID {
	u.mu.Lock()
	defer u.mu.Unlock()

	receiptID := uuid.New()
	receiptPoints := int64(0)

	// Rule 1. One point for every alphanumeric character in the retailer name.
	retailerMatches := alphanumericRegexp.FindAllString(receipt.Retailer, -1)
	receiptPoints += int64(len(retailerMatches))

	// Rule 2. 50 points if the total is a round dollar amount with no cents.
	if math.Mod(receipt.Total, 1) == 0 {
		receiptPoints += int64(50)
	}

	// Rule 3. 25 points if the total is a multiple of `0.25`
	if math.Mod(receipt.Total, 0.25) == 0 {
		receiptPoints += 25
	}

	// Rule 4. 5 points for every two items on the receipt.
	receiptPoints += int64(len(receipt.Items)/2) * 5

	// Rule 5. If the trimmed length of the item description is a multiple of 3, multiply the price by `0.2` and round up to the nearest integer. The result is the number of points earned.
	for _, item := range receipt.Items {
		trimmedLength := float64(len(strings.TrimSpace(item.ShortDescription)))

		if math.Mod(trimmedLength, 3) == 0 {
			itemPoints := math.Ceil(item.Price * 0.2)
			receiptPoints += int64(itemPoints)
		}
	}

	// Rule 6. 6 points if the day in the purchase date is odd.
	purchaseDay := float64(receipt.PurchasedAt.Day())
	if math.Mod(purchaseDay, 2) != 0 {
		receiptPoints += 6
	}

	// Rule 7. 10 points if the time of purchase is after 2:00pm and before 4:00pm
	// Encode the purchase time as a float so I can do a simple less than and greater than check
	purchaseFloat := float64(receipt.PurchasedAt.Hour()) + (float64(receipt.PurchasedAt.Minute()) / 60)

	// 'm interpreting that the after 2 and before 4 doesn't include 2:00 and 4:00 respectively
	if purchaseFloat > 14 && purchaseFloat < 16 {
		receiptPoints += 10
	}

	u.points[receiptID] = receiptPoints
	return receiptID
}

func (u *Usecases) GetReceiptPoints(ctx context.Context, id uuid.UUID) (int64, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	points, ok := u.points[id]
	if !ok {
		return 0, fmt.Errorf("receipt could not be found")
	}

	return points, nil
}
