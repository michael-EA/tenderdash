// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	bytes "github.com/tendermint/tendermint/libs/bytes"
	client "github.com/tendermint/tendermint/rpc/client"

	context "context"

	coretypes "github.com/tendermint/tendermint/rpc/coretypes"

	mock "github.com/stretchr/testify/mock"

	testing "testing"

	types "github.com/tendermint/tendermint/types"
)

// ABCIClient is an autogenerated mock type for the ABCIClient type
type ABCIClient struct {
	mock.Mock
}

// ABCIInfo provides a mock function with given fields: _a0
func (_m *ABCIClient) ABCIInfo(_a0 context.Context) (*coretypes.ResultABCIInfo, error) {
	ret := _m.Called(_a0)

	var r0 *coretypes.ResultABCIInfo
	if rf, ok := ret.Get(0).(func(context.Context) *coretypes.ResultABCIInfo); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.ResultABCIInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ABCIQuery provides a mock function with given fields: ctx, path, data
func (_m *ABCIClient) ABCIQuery(ctx context.Context, path string, data bytes.HexBytes) (*coretypes.ResultABCIQuery, error) {
	ret := _m.Called(ctx, path, data)

	var r0 *coretypes.ResultABCIQuery
	if rf, ok := ret.Get(0).(func(context.Context, string, bytes.HexBytes) *coretypes.ResultABCIQuery); ok {
		r0 = rf(ctx, path, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.ResultABCIQuery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, bytes.HexBytes) error); ok {
		r1 = rf(ctx, path, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ABCIQueryWithOptions provides a mock function with given fields: ctx, path, data, opts
func (_m *ABCIClient) ABCIQueryWithOptions(ctx context.Context, path string, data bytes.HexBytes, opts client.ABCIQueryOptions) (*coretypes.ResultABCIQuery, error) {
	ret := _m.Called(ctx, path, data, opts)

	var r0 *coretypes.ResultABCIQuery
	if rf, ok := ret.Get(0).(func(context.Context, string, bytes.HexBytes, client.ABCIQueryOptions) *coretypes.ResultABCIQuery); ok {
		r0 = rf(ctx, path, data, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.ResultABCIQuery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, bytes.HexBytes, client.ABCIQueryOptions) error); ok {
		r1 = rf(ctx, path, data, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BroadcastTxAsync provides a mock function with given fields: _a0, _a1
func (_m *ABCIClient) BroadcastTxAsync(_a0 context.Context, _a1 types.Tx) (*coretypes.ResultBroadcastTx, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *coretypes.ResultBroadcastTx
	if rf, ok := ret.Get(0).(func(context.Context, types.Tx) *coretypes.ResultBroadcastTx); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.ResultBroadcastTx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.Tx) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BroadcastTxCommit provides a mock function with given fields: _a0, _a1
func (_m *ABCIClient) BroadcastTxCommit(_a0 context.Context, _a1 types.Tx) (*coretypes.ResultBroadcastTxCommit, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *coretypes.ResultBroadcastTxCommit
	if rf, ok := ret.Get(0).(func(context.Context, types.Tx) *coretypes.ResultBroadcastTxCommit); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.ResultBroadcastTxCommit)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.Tx) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BroadcastTxSync provides a mock function with given fields: _a0, _a1
func (_m *ABCIClient) BroadcastTxSync(_a0 context.Context, _a1 types.Tx) (*coretypes.ResultBroadcastTx, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *coretypes.ResultBroadcastTx
	if rf, ok := ret.Get(0).(func(context.Context, types.Tx) *coretypes.ResultBroadcastTx); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.ResultBroadcastTx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.Tx) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewABCIClient creates a new instance of ABCIClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewABCIClient(t testing.TB) *ABCIClient {
	mock := &ABCIClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
