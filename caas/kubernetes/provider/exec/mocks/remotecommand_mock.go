// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/tools/remotecommand (interfaces: Executor)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/remotecommand_mock.go k8s.io/client-go/tools/remotecommand Executor
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	remotecommand "k8s.io/client-go/tools/remotecommand"
)

// MockExecutor is a mock of Executor interface.
type MockExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockExecutorMockRecorder
}

// MockExecutorMockRecorder is the mock recorder for MockExecutor.
type MockExecutorMockRecorder struct {
	mock *MockExecutor
}

// NewMockExecutor creates a new mock instance.
func NewMockExecutor(ctrl *gomock.Controller) *MockExecutor {
	mock := &MockExecutor{ctrl: ctrl}
	mock.recorder = &MockExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecutor) EXPECT() *MockExecutorMockRecorder {
	return m.recorder
}

// Stream mocks base method.
func (m *MockExecutor) Stream(arg0 remotecommand.StreamOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stream", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stream indicates an expected call of Stream.
func (mr *MockExecutorMockRecorder) Stream(arg0 any) *MockExecutorStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stream", reflect.TypeOf((*MockExecutor)(nil).Stream), arg0)
	return &MockExecutorStreamCall{Call: call}
}

// MockExecutorStreamCall wrap *gomock.Call
type MockExecutorStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExecutorStreamCall) Return(arg0 error) *MockExecutorStreamCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExecutorStreamCall) Do(f func(remotecommand.StreamOptions) error) *MockExecutorStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExecutorStreamCall) DoAndReturn(f func(remotecommand.StreamOptions) error) *MockExecutorStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StreamWithContext mocks base method.
func (m *MockExecutor) StreamWithContext(arg0 context.Context, arg1 remotecommand.StreamOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamWithContext", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StreamWithContext indicates an expected call of StreamWithContext.
func (mr *MockExecutorMockRecorder) StreamWithContext(arg0, arg1 any) *MockExecutorStreamWithContextCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamWithContext", reflect.TypeOf((*MockExecutor)(nil).StreamWithContext), arg0, arg1)
	return &MockExecutorStreamWithContextCall{Call: call}
}

// MockExecutorStreamWithContextCall wrap *gomock.Call
type MockExecutorStreamWithContextCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExecutorStreamWithContextCall) Return(arg0 error) *MockExecutorStreamWithContextCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExecutorStreamWithContextCall) Do(f func(context.Context, remotecommand.StreamOptions) error) *MockExecutorStreamWithContextCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExecutorStreamWithContextCall) DoAndReturn(f func(context.Context, remotecommand.StreamOptions) error) *MockExecutorStreamWithContextCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
