// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/worker/gate (interfaces: Lock)
//
// Generated by this command:
//
//	mockgen -typed -package upgradedatabase -destination lock_mock_test.go github.com/juju/juju/internal/worker/gate Lock
//

// Package upgradedatabase is a generated GoMock package.
package upgradedatabase

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockLock is a mock of Lock interface.
type MockLock struct {
	ctrl     *gomock.Controller
	recorder *MockLockMockRecorder
}

// MockLockMockRecorder is the mock recorder for MockLock.
type MockLockMockRecorder struct {
	mock *MockLock
}

// NewMockLock creates a new mock instance.
func NewMockLock(ctrl *gomock.Controller) *MockLock {
	mock := &MockLock{ctrl: ctrl}
	mock.recorder = &MockLockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLock) EXPECT() *MockLockMockRecorder {
	return m.recorder
}

// IsUnlocked mocks base method.
func (m *MockLock) IsUnlocked() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUnlocked")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUnlocked indicates an expected call of IsUnlocked.
func (mr *MockLockMockRecorder) IsUnlocked() *MockLockIsUnlockedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUnlocked", reflect.TypeOf((*MockLock)(nil).IsUnlocked))
	return &MockLockIsUnlockedCall{Call: call}
}

// MockLockIsUnlockedCall wrap *gomock.Call
type MockLockIsUnlockedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLockIsUnlockedCall) Return(arg0 bool) *MockLockIsUnlockedCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLockIsUnlockedCall) Do(f func() bool) *MockLockIsUnlockedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLockIsUnlockedCall) DoAndReturn(f func() bool) *MockLockIsUnlockedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Unlock mocks base method.
func (m *MockLock) Unlock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unlock")
}

// Unlock indicates an expected call of Unlock.
func (mr *MockLockMockRecorder) Unlock() *MockLockUnlockCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockLock)(nil).Unlock))
	return &MockLockUnlockCall{Call: call}
}

// MockLockUnlockCall wrap *gomock.Call
type MockLockUnlockCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLockUnlockCall) Return() *MockLockUnlockCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLockUnlockCall) Do(f func()) *MockLockUnlockCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLockUnlockCall) DoAndReturn(f func()) *MockLockUnlockCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Unlocked mocks base method.
func (m *MockLock) Unlocked() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unlocked")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Unlocked indicates an expected call of Unlocked.
func (mr *MockLockMockRecorder) Unlocked() *MockLockUnlockedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlocked", reflect.TypeOf((*MockLock)(nil).Unlocked))
	return &MockLockUnlockedCall{Call: call}
}

// MockLockUnlockedCall wrap *gomock.Call
type MockLockUnlockedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLockUnlockedCall) Return(arg0 <-chan struct{}) *MockLockUnlockedCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLockUnlockedCall) Do(f func() <-chan struct{}) *MockLockUnlockedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLockUnlockedCall) DoAndReturn(f func() <-chan struct{}) *MockLockUnlockedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
