package http

import (
	"context"

	"receipt-processor-challenge/internal/delivery/oapi"

	"github.com/google/uuid"
)

// Submits a receipt for processing.
// (POST /receipts/process)
func (repo *HTTPRepo) PostReceiptsProcess(ctx context.Context, request oapi.PostReceiptsProcessRequestObject) (oapi.PostReceiptsProcessResponseObject, error) {
	if request.Body == nil {
		return oapi.PostReceiptsProcess400Response{}, nil
	}

	receiptDto, err := oapi.Receipt(*request.Body).ToDTO()
	if err != nil {
		return oapi.PostReceiptsProcess400Response{}, nil
	}

	id := repo.UsecasesRepo.ProcessReceipt(ctx, receiptDto)
	return oapi.PostReceiptsProcess200JSONResponse{Id: id.String()}, nil
}

// Returns the points awarded for the receipt.
// (GET /receipts/{id}/points)
func (repo *HTTPRepo) GetReceiptsIdPoints(ctx context.Context, request oapi.GetReceiptsIdPointsRequestObject) (oapi.GetReceiptsIdPointsResponseObject, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return oapi.GetReceiptsIdPoints404Response{}, nil
	}

	points, err := repo.UsecasesRepo.GetReceiptPoints(ctx, id)
	if err != nil {
		return oapi.GetReceiptsIdPoints404Response{}, nil
	}

	return oapi.GetReceiptsIdPoints200JSONResponse{Points: &points}, nil
}

// Deletes a receipt
// (DELETE /receipts/{id})
func (repo *HTTPRepo) DeleteReceiptsId(ctx context.Context, request oapi.DeleteReceiptsIdRequestObject) (oapi.DeleteReceiptsIdResponseObject, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return oapi.GetReceiptsIdPoints404Response{}, nil
	}

	err = repo.UsecasesRepo.DeleteReceipt(ctx, id)
	if err != nil {
		return oapi.DeleteReceiptsId404Response{}, err
	}

	return oapi.DeleteReceiptsId204Response{}, nil
}
