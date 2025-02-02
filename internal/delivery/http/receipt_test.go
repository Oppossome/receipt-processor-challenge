package http_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"receipt-processor-challenge/internal/delivery/http"
	"receipt-processor-challenge/internal/delivery/oapi"
	mock_usecases "receipt-processor-challenge/internal/mocks/domain/usecases"

	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func NewTestServer(t *testing.T) (*httptest.Server, *mock_usecases.MockUsecasesRepo) {
	// Setup our mock
	mockUsecases := mock_usecases.NewMockUsecasesRepo(t)
	testHTTP := http.HTTPRepo{UsecasesRepo: mockUsecases}

	chiRouter, err := oapi.NewChiRouter(&testHTTP)
	assert.NoError(t, err)

	server := httptest.NewServer(chiRouter)
	return server, mockUsecases
}

func TestPostReceiptsProcess(t *testing.T) {
	t.Run("400 - Invalid Receipt", func(t *testing.T) {
		testServer, _ := NewTestServer(t)
		testClient := testServer.Client()
		defer testServer.Close()

		request, err := testClient.Post(
			fmt.Sprintf("%s/receipts/process", testServer.URL),
			"application/json",
			bytes.NewBufferString("{}"),
		)

		assert.NoError(t, err)
		assert.Equal(t, 400, request.StatusCode)
	})

	t.Run("200 - Ok", func(t *testing.T) {
		testServer, mocks := NewTestServer(t)
		testClient := testServer.Client()
		defer testServer.Close()

		// Setup our expected input and output
		oapiReceipt := oapi.Receipt{
			Retailer:     "M&M Corner Market",
			PurchaseDate: openapi_types.Date{Time: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local)},
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
		}

		dtoReceipt, err := oapiReceipt.ToDTO()
		assert.NoError(t, err)

		// Configure the mock to expect our dtoReceipt
		receiptUUID := uuid.New()
		mocks.EXPECT().ProcessReceipt(mock.Anything, dtoReceipt).Return(receiptUUID)

		// Turn our request receipt into
		jsonBytes, err := json.Marshal(oapiReceipt)
		assert.NoError(t, err)

		// Perform the request.
		request, err := testClient.Post(
			fmt.Sprintf("%s/receipts/process", testServer.URL),
			"application/json",
			bytes.NewBuffer(jsonBytes),
		)

		assert.NoError(t, err)
		assert.Equal(t, 200, request.StatusCode)

		// Process our request
		var response *oapi.PostReceiptsProcess200JSONResponse
		err = json.NewDecoder(request.Body).Decode(&response)
		assert.NoError(t, err)

		assert.Equal(t, response.Id, receiptUUID.String())
	})
}

func TestGetReceiptsIdPoints(t *testing.T) {
	t.Run("404 - Invalid Receipt Id", func(t *testing.T) {
		testServer, _ := NewTestServer(t)
		testClient := testServer.Client()
		defer testServer.Close()

		request, err := testClient.Get(fmt.Sprintf("%s/receipts/123/points", testServer.URL))

		assert.NoError(t, err)
		assert.Equal(t, 404, request.StatusCode)
	})

	t.Run("404 - Receipt not found", func(t *testing.T) {
		testServer, mocks := NewTestServer(t)
		testClient := testServer.Client()
		defer testServer.Close()

		receiptUUID := uuid.New()
		mocks.EXPECT().GetReceiptPoints(mock.Anything, receiptUUID).Return(0, fmt.Errorf("receipt could not be found"))

		request, err := testClient.Get(fmt.Sprintf("%s/receipts/%s/points", testServer.URL, receiptUUID))

		assert.NoError(t, err)
		assert.Equal(t, 404, request.StatusCode)
	})

	t.Run("200 - Ok", func(t *testing.T) {
		testServer, mocks := NewTestServer(t)
		testClient := testServer.Client()
		defer testServer.Close()

		receiptUUID := uuid.New()
		receiptPoints := int64(123)

		mocks.EXPECT().GetReceiptPoints(mock.Anything, receiptUUID).Return(receiptPoints, nil)
		request, err := testClient.Get(fmt.Sprintf("%s/receipts/%s/points", testServer.URL, receiptUUID))

		assert.NoError(t, err)
		assert.Equal(t, 200, request.StatusCode)

		var response *oapi.GetReceiptsIdPoints200JSONResponse
		err = json.NewDecoder(request.Body).Decode(&response)
		assert.NoError(t, err)

		assert.Equal(t, receiptPoints, *response.Points)
	})
}
