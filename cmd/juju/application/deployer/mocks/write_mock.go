// Code generated by MockGen. DO NOT EDIT.
// Source: io (interfaces: Writer)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/write_mock.go io Writer
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockWriter is a mock of Writer interface.
type MockWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWriterMockRecorder
}

// MockWriterMockRecorder is the mock recorder for MockWriter.
type MockWriterMockRecorder struct {
	mock *MockWriter
}

// NewMockWriter creates a new mock instance.
func NewMockWriter(ctrl *gomock.Controller) *MockWriter {
	mock := &MockWriter{ctrl: ctrl}
	mock.recorder = &MockWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWriter) EXPECT() *MockWriterMockRecorder {
	return m.recorder
}

// Write mocks base method.
func (m *MockWriter) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockWriterMockRecorder) Write(arg0 any) *MockWriterWriteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockWriter)(nil).Write), arg0)
	return &MockWriterWriteCall{Call: call}
}

// MockWriterWriteCall wrap *gomock.Call
type MockWriterWriteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockWriterWriteCall) Return(arg0 int, arg1 error) *MockWriterWriteCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockWriterWriteCall) Do(f func([]byte) (int, error)) *MockWriterWriteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockWriterWriteCall) DoAndReturn(f func([]byte) (int, error)) *MockWriterWriteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
