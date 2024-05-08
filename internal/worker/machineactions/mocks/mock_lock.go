// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/machinelock (interfaces: Lock)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/mock_lock.go github.com/juju/juju/core/machinelock Lock
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	machinelock "github.com/juju/juju/core/machinelock"
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

// Acquire mocks base method.
func (m *MockLock) Acquire(arg0 machinelock.Spec) (func(), error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Acquire", arg0)
	ret0, _ := ret[0].(func())
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Acquire indicates an expected call of Acquire.
func (mr *MockLockMockRecorder) Acquire(arg0 any) *MockLockAcquireCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Acquire", reflect.TypeOf((*MockLock)(nil).Acquire), arg0)
	return &MockLockAcquireCall{Call: call}
}

// MockLockAcquireCall wrap *gomock.Call
type MockLockAcquireCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLockAcquireCall) Return(arg0 func(), arg1 error) *MockLockAcquireCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLockAcquireCall) Do(f func(machinelock.Spec) (func(), error)) *MockLockAcquireCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLockAcquireCall) DoAndReturn(f func(machinelock.Spec) (func(), error)) *MockLockAcquireCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Report mocks base method.
func (m *MockLock) Report(arg0 ...machinelock.ReportOption) (string, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Report", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Report indicates an expected call of Report.
func (mr *MockLockMockRecorder) Report(arg0 ...any) *MockLockReportCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Report", reflect.TypeOf((*MockLock)(nil).Report), arg0...)
	return &MockLockReportCall{Call: call}
}

// MockLockReportCall wrap *gomock.Call
type MockLockReportCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLockReportCall) Return(arg0 string, arg1 error) *MockLockReportCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLockReportCall) Do(f func(...machinelock.ReportOption) (string, error)) *MockLockReportCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLockReportCall) DoAndReturn(f func(...machinelock.ReportOption) (string, error)) *MockLockReportCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
