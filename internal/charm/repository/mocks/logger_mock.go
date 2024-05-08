// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/charm/repository (interfaces: Logger)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/logger_mock.go github.com/juju/juju/internal/charm/repository Logger
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Errorf mocks base method.
func (m *MockLogger) Errorf(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockLoggerMockRecorder) Errorf(arg0 any, arg1 ...any) *MockLoggerErrorfCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLogger)(nil).Errorf), varargs...)
	return &MockLoggerErrorfCall{Call: call}
}

// MockLoggerErrorfCall wrap *gomock.Call
type MockLoggerErrorfCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoggerErrorfCall) Return() *MockLoggerErrorfCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoggerErrorfCall) Do(f func(string, ...any)) *MockLoggerErrorfCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoggerErrorfCall) DoAndReturn(f func(string, ...any)) *MockLoggerErrorfCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Tracef mocks base method.
func (m *MockLogger) Tracef(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Tracef", varargs...)
}

// Tracef indicates an expected call of Tracef.
func (mr *MockLoggerMockRecorder) Tracef(arg0 any, arg1 ...any) *MockLoggerTracefCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tracef", reflect.TypeOf((*MockLogger)(nil).Tracef), varargs...)
	return &MockLoggerTracefCall{Call: call}
}

// MockLoggerTracefCall wrap *gomock.Call
type MockLoggerTracefCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoggerTracefCall) Return() *MockLoggerTracefCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoggerTracefCall) Do(f func(string, ...any)) *MockLoggerTracefCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoggerTracefCall) DoAndReturn(f func(string, ...any)) *MockLoggerTracefCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
