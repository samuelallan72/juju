// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/controller/crossmodelsecrets (interfaces: CrossModelState)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/crossmodel.go github.com/juju/juju/apiserver/facades/controller/crossmodelsecrets CrossModelState
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	names "github.com/juju/names/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockCrossModelState is a mock of CrossModelState interface.
type MockCrossModelState struct {
	ctrl     *gomock.Controller
	recorder *MockCrossModelStateMockRecorder
}

// MockCrossModelStateMockRecorder is the mock recorder for MockCrossModelState.
type MockCrossModelStateMockRecorder struct {
	mock *MockCrossModelState
}

// NewMockCrossModelState creates a new mock instance.
func NewMockCrossModelState(ctrl *gomock.Controller) *MockCrossModelState {
	mock := &MockCrossModelState{ctrl: ctrl}
	mock.recorder = &MockCrossModelStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCrossModelState) EXPECT() *MockCrossModelStateMockRecorder {
	return m.recorder
}

// GetRemoteApplicationTag mocks base method.
func (m *MockCrossModelState) GetRemoteApplicationTag(arg0 string) (names.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemoteApplicationTag", arg0)
	ret0, _ := ret[0].(names.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRemoteApplicationTag indicates an expected call of GetRemoteApplicationTag.
func (mr *MockCrossModelStateMockRecorder) GetRemoteApplicationTag(arg0 any) *MockCrossModelStateGetRemoteApplicationTagCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteApplicationTag", reflect.TypeOf((*MockCrossModelState)(nil).GetRemoteApplicationTag), arg0)
	return &MockCrossModelStateGetRemoteApplicationTagCall{Call: call}
}

// MockCrossModelStateGetRemoteApplicationTagCall wrap *gomock.Call
type MockCrossModelStateGetRemoteApplicationTagCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCrossModelStateGetRemoteApplicationTagCall) Return(arg0 names.Tag, arg1 error) *MockCrossModelStateGetRemoteApplicationTagCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCrossModelStateGetRemoteApplicationTagCall) Do(f func(string) (names.Tag, error)) *MockCrossModelStateGetRemoteApplicationTagCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCrossModelStateGetRemoteApplicationTagCall) DoAndReturn(f func(string) (names.Tag, error)) *MockCrossModelStateGetRemoteApplicationTagCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetToken mocks base method.
func (m *MockCrossModelState) GetToken(arg0 names.Tag) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToken", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToken indicates an expected call of GetToken.
func (mr *MockCrossModelStateMockRecorder) GetToken(arg0 any) *MockCrossModelStateGetTokenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken", reflect.TypeOf((*MockCrossModelState)(nil).GetToken), arg0)
	return &MockCrossModelStateGetTokenCall{Call: call}
}

// MockCrossModelStateGetTokenCall wrap *gomock.Call
type MockCrossModelStateGetTokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCrossModelStateGetTokenCall) Return(arg0 string, arg1 error) *MockCrossModelStateGetTokenCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCrossModelStateGetTokenCall) Do(f func(names.Tag) (string, error)) *MockCrossModelStateGetTokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCrossModelStateGetTokenCall) DoAndReturn(f func(names.Tag) (string, error)) *MockCrossModelStateGetTokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
