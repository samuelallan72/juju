// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/lease (interfaces: ExpiryStore)
//
// Generated by this command:
//
//	mockgen -typed -package leaseexpiry_test -destination store_mock_test.go github.com/juju/juju/core/lease ExpiryStore
//

// Package leaseexpiry_test is a generated GoMock package.
package leaseexpiry_test

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockExpiryStore is a mock of ExpiryStore interface.
type MockExpiryStore struct {
	ctrl     *gomock.Controller
	recorder *MockExpiryStoreMockRecorder
}

// MockExpiryStoreMockRecorder is the mock recorder for MockExpiryStore.
type MockExpiryStoreMockRecorder struct {
	mock *MockExpiryStore
}

// NewMockExpiryStore creates a new mock instance.
func NewMockExpiryStore(ctrl *gomock.Controller) *MockExpiryStore {
	mock := &MockExpiryStore{ctrl: ctrl}
	mock.recorder = &MockExpiryStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExpiryStore) EXPECT() *MockExpiryStoreMockRecorder {
	return m.recorder
}

// ExpireLeases mocks base method.
func (m *MockExpiryStore) ExpireLeases(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpireLeases", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExpireLeases indicates an expected call of ExpireLeases.
func (mr *MockExpiryStoreMockRecorder) ExpireLeases(arg0 any) *MockExpiryStoreExpireLeasesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpireLeases", reflect.TypeOf((*MockExpiryStore)(nil).ExpireLeases), arg0)
	return &MockExpiryStoreExpireLeasesCall{Call: call}
}

// MockExpiryStoreExpireLeasesCall wrap *gomock.Call
type MockExpiryStoreExpireLeasesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExpiryStoreExpireLeasesCall) Return(arg0 error) *MockExpiryStoreExpireLeasesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExpiryStoreExpireLeasesCall) Do(f func(context.Context) error) *MockExpiryStoreExpireLeasesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExpiryStoreExpireLeasesCall) DoAndReturn(f func(context.Context) error) *MockExpiryStoreExpireLeasesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
