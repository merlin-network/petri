// Code generated by MockGen. DO NOT EDIT.
// Source: x/interchainqueries/types/verify.go

// Package mock_types is a generated GoMock package.
package mock_types

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/codec/types"
	types0 "github.com/cosmos/cosmos-sdk/types"
	keeper "github.com/cosmos/ibc-go/v4/modules/core/02-client/keeper"
	exported "github.com/cosmos/ibc-go/v4/modules/core/exported"
	types1 "github.com/cosmos/ibc-go/v4/modules/light-clients/07-tendermint/types"
	gomock "github.com/golang/mock/gomock"

	types2 "github.com/petri-labs/petri/x/interchainqueries/types"
)

// MockHeaderVerifier is a mock of HeaderVerifier interface.
type MockHeaderVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockHeaderVerifierMockRecorder
}

// MockHeaderVerifierMockRecorder is the mock recorder for MockHeaderVerifier.
type MockHeaderVerifierMockRecorder struct {
	mock *MockHeaderVerifier
}

// NewMockHeaderVerifier creates a new mock instance.
func NewMockHeaderVerifier(ctrl *gomock.Controller) *MockHeaderVerifier {
	mock := &MockHeaderVerifier{ctrl: ctrl}
	mock.recorder = &MockHeaderVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHeaderVerifier) EXPECT() *MockHeaderVerifierMockRecorder {
	return m.recorder
}

// UnpackHeader mocks base method.
func (m *MockHeaderVerifier) UnpackHeader(any *types.Any) (exported.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnpackHeader", any)
	ret0, _ := ret[0].(exported.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnpackHeader indicates an expected call of UnpackHeader.
func (mr *MockHeaderVerifierMockRecorder) UnpackHeader(any interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpackHeader", reflect.TypeOf((*MockHeaderVerifier)(nil).UnpackHeader), any)
}

// VerifyHeaders mocks base method.
func (m *MockHeaderVerifier) VerifyHeaders(ctx types0.Context, cleintkeeper keeper.Keeper, clientID string, header, nextHeader exported.Header) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyHeaders", ctx, cleintkeeper, clientID, header, nextHeader)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyHeaders indicates an expected call of VerifyHeaders.
func (mr *MockHeaderVerifierMockRecorder) VerifyHeaders(ctx, cleintkeeper, clientID, header, nextHeader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyHeaders", reflect.TypeOf((*MockHeaderVerifier)(nil).VerifyHeaders), ctx, cleintkeeper, clientID, header, nextHeader)
}

// MockTransactionVerifier is a mock of TransactionVerifier interface.
type MockTransactionVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionVerifierMockRecorder
}

// MockTransactionVerifierMockRecorder is the mock recorder for MockTransactionVerifier.
type MockTransactionVerifierMockRecorder struct {
	mock *MockTransactionVerifier
}

// NewMockTransactionVerifier creates a new mock instance.
func NewMockTransactionVerifier(ctrl *gomock.Controller) *MockTransactionVerifier {
	mock := &MockTransactionVerifier{ctrl: ctrl}
	mock.recorder = &MockTransactionVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionVerifier) EXPECT() *MockTransactionVerifierMockRecorder {
	return m.recorder
}

// VerifyTransaction mocks base method.
func (m *MockTransactionVerifier) VerifyTransaction(header, nextHeader *types1.Header, tx *types2.TxValue) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyTransaction", header, nextHeader, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyTransaction indicates an expected call of VerifyTransaction.
func (mr *MockTransactionVerifierMockRecorder) VerifyTransaction(header, nextHeader, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyTransaction", reflect.TypeOf((*MockTransactionVerifier)(nil).VerifyTransaction), header, nextHeader, tx)
}
