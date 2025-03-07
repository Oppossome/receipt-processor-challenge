// Code generated by mockery v2.52.1. DO NOT EDIT.

package mock_usecases

import (
	context "context"
	domain "receipt-processor-challenge/internal/domain"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockUsecasesRepo is an autogenerated mock type for the UsecasesRepo type
type MockUsecasesRepo struct {
	mock.Mock
}

type MockUsecasesRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUsecasesRepo) EXPECT() *MockUsecasesRepo_Expecter {
	return &MockUsecasesRepo_Expecter{mock: &_m.Mock}
}

// DeleteReceipt provides a mock function with given fields: ctx, id
func (_m *MockUsecasesRepo) DeleteReceipt(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteReceipt")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUsecasesRepo_DeleteReceipt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteReceipt'
type MockUsecasesRepo_DeleteReceipt_Call struct {
	*mock.Call
}

// DeleteReceipt is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *MockUsecasesRepo_Expecter) DeleteReceipt(ctx interface{}, id interface{}) *MockUsecasesRepo_DeleteReceipt_Call {
	return &MockUsecasesRepo_DeleteReceipt_Call{Call: _e.mock.On("DeleteReceipt", ctx, id)}
}

func (_c *MockUsecasesRepo_DeleteReceipt_Call) Run(run func(ctx context.Context, id uuid.UUID)) *MockUsecasesRepo_DeleteReceipt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *MockUsecasesRepo_DeleteReceipt_Call) Return(_a0 error) *MockUsecasesRepo_DeleteReceipt_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUsecasesRepo_DeleteReceipt_Call) RunAndReturn(run func(context.Context, uuid.UUID) error) *MockUsecasesRepo_DeleteReceipt_Call {
	_c.Call.Return(run)
	return _c
}

// GetReceiptPoints provides a mock function with given fields: ctx, id
func (_m *MockUsecasesRepo) GetReceiptPoints(ctx context.Context, id uuid.UUID) (int64, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetReceiptPoints")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (int64, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) int64); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUsecasesRepo_GetReceiptPoints_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetReceiptPoints'
type MockUsecasesRepo_GetReceiptPoints_Call struct {
	*mock.Call
}

// GetReceiptPoints is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *MockUsecasesRepo_Expecter) GetReceiptPoints(ctx interface{}, id interface{}) *MockUsecasesRepo_GetReceiptPoints_Call {
	return &MockUsecasesRepo_GetReceiptPoints_Call{Call: _e.mock.On("GetReceiptPoints", ctx, id)}
}

func (_c *MockUsecasesRepo_GetReceiptPoints_Call) Run(run func(ctx context.Context, id uuid.UUID)) *MockUsecasesRepo_GetReceiptPoints_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *MockUsecasesRepo_GetReceiptPoints_Call) Return(_a0 int64, _a1 error) *MockUsecasesRepo_GetReceiptPoints_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUsecasesRepo_GetReceiptPoints_Call) RunAndReturn(run func(context.Context, uuid.UUID) (int64, error)) *MockUsecasesRepo_GetReceiptPoints_Call {
	_c.Call.Return(run)
	return _c
}

// ProcessReceipt provides a mock function with given fields: ctx, receipt
func (_m *MockUsecasesRepo) ProcessReceipt(ctx context.Context, receipt domain.ReceiptDTO) uuid.UUID {
	ret := _m.Called(ctx, receipt)

	if len(ret) == 0 {
		panic("no return value specified for ProcessReceipt")
	}

	var r0 uuid.UUID
	if rf, ok := ret.Get(0).(func(context.Context, domain.ReceiptDTO) uuid.UUID); ok {
		r0 = rf(ctx, receipt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	return r0
}

// MockUsecasesRepo_ProcessReceipt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessReceipt'
type MockUsecasesRepo_ProcessReceipt_Call struct {
	*mock.Call
}

// ProcessReceipt is a helper method to define mock.On call
//   - ctx context.Context
//   - receipt domain.ReceiptDTO
func (_e *MockUsecasesRepo_Expecter) ProcessReceipt(ctx interface{}, receipt interface{}) *MockUsecasesRepo_ProcessReceipt_Call {
	return &MockUsecasesRepo_ProcessReceipt_Call{Call: _e.mock.On("ProcessReceipt", ctx, receipt)}
}

func (_c *MockUsecasesRepo_ProcessReceipt_Call) Run(run func(ctx context.Context, receipt domain.ReceiptDTO)) *MockUsecasesRepo_ProcessReceipt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.ReceiptDTO))
	})
	return _c
}

func (_c *MockUsecasesRepo_ProcessReceipt_Call) Return(_a0 uuid.UUID) *MockUsecasesRepo_ProcessReceipt_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUsecasesRepo_ProcessReceipt_Call) RunAndReturn(run func(context.Context, domain.ReceiptDTO) uuid.UUID) *MockUsecasesRepo_ProcessReceipt_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUsecasesRepo creates a new instance of MockUsecasesRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUsecasesRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUsecasesRepo {
	mock := &MockUsecasesRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
