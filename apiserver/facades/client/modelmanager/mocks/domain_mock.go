// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/common (interfaces: ControllerConfigService)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/domain_mock.go github.com/juju/juju/apiserver/common ControllerConfigService
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	controller "github.com/juju/juju/controller"
	gomock "go.uber.org/mock/gomock"
)

// MockControllerConfigService is a mock of ControllerConfigService interface.
type MockControllerConfigService struct {
	ctrl     *gomock.Controller
	recorder *MockControllerConfigServiceMockRecorder
}

// MockControllerConfigServiceMockRecorder is the mock recorder for MockControllerConfigService.
type MockControllerConfigServiceMockRecorder struct {
	mock *MockControllerConfigService
}

// NewMockControllerConfigService creates a new mock instance.
func NewMockControllerConfigService(ctrl *gomock.Controller) *MockControllerConfigService {
	mock := &MockControllerConfigService{ctrl: ctrl}
	mock.recorder = &MockControllerConfigServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControllerConfigService) EXPECT() *MockControllerConfigServiceMockRecorder {
	return m.recorder
}

// ControllerConfig mocks base method.
func (m *MockControllerConfigService) ControllerConfig(arg0 context.Context) (controller.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig", arg0)
	ret0, _ := ret[0].(controller.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockControllerConfigServiceMockRecorder) ControllerConfig(arg0 any) *MockControllerConfigServiceControllerConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockControllerConfigService)(nil).ControllerConfig), arg0)
	return &MockControllerConfigServiceControllerConfigCall{Call: call}
}

// MockControllerConfigServiceControllerConfigCall wrap *gomock.Call
type MockControllerConfigServiceControllerConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerConfigServiceControllerConfigCall) Return(arg0 controller.Config, arg1 error) *MockControllerConfigServiceControllerConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerConfigServiceControllerConfigCall) Do(f func(context.Context) (controller.Config, error)) *MockControllerConfigServiceControllerConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerConfigServiceControllerConfigCall) DoAndReturn(f func(context.Context) (controller.Config, error)) *MockControllerConfigServiceControllerConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
