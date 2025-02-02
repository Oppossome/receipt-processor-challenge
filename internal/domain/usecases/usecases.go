package usecases

import (
	"context"
	"sync"

	"receipt-processor-challenge/internal/domain"

	"github.com/google/uuid"
)

type UsecasesRepo interface {
	ProcessReceipt(ctx context.Context, receipt domain.ReceiptDTO) uuid.UUID
	GetReceiptPoints(ctx context.Context, id uuid.UUID) (int64, error)
}

// Since circumstances don't demand we have a database, I'll go with the simplest solution for now
type Usecases struct {
	mu     sync.Mutex
	points map[uuid.UUID]int64
}

// Check that we conform to ServerInterface
var _ UsecasesRepo = (*Usecases)(nil)

func NewUsecases() *Usecases {
	return &Usecases{
		points: make(map[uuid.UUID]int64),
	}
}
