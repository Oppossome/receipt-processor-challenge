package http

import (
	"receipt-processor-challenge/internal/delivery/oapi"
	"receipt-processor-challenge/internal/domain/usecases"
)

type HTTPRepo struct {
	UsecasesRepo usecases.UsecasesRepo
}

// Check that we conform to ServerInterface
var _ oapi.StrictServerInterface = (*HTTPRepo)(nil)
