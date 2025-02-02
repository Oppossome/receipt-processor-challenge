package oapi

import (
	"strconv"
	"time"

	"receipt-processor-challenge/internal/domain"
)

func (item Item) ToDTO() (domain.ItemDTO, error) {
	price, err := strconv.ParseFloat(item.Price, 64)
	if err != nil {
		return domain.ItemDTO{}, err
	}

	return domain.ItemDTO{
		ShortDescription: item.ShortDescription,
		Price:            price,
	}, nil
}

func (receipt Receipt) ToDTO() (domain.ReceiptDTO, error) {
	items := make([]domain.ItemDTO, len(receipt.Items))
	for i, item := range receipt.Items {
		dtoItem, err := item.ToDTO()
		if err != nil {
			return domain.ReceiptDTO{}, err
		}
		items[i] = dtoItem
	}

	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return domain.ReceiptDTO{}, err
	}

	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		return domain.ReceiptDTO{}, err
	}

	return domain.ReceiptDTO{
		Items: items,
		// There's probably a better way to do this but according to my tests this works 😎
		PurchasedAt: time.Date(
			receipt.PurchaseDate.Year(),
			receipt.PurchaseDate.Month(),
			receipt.PurchaseDate.Day(),
			purchaseTime.Hour(),
			purchaseTime.Minute(),
			0, 0, time.Local,
		),
		Retailer: receipt.Retailer,
		Total:    total,
	}, nil
}
