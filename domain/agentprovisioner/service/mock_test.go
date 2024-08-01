// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/agentprovisioner/service (interfaces: State,Provider)
//
// Generated by this command:
//
//	mockgen -typed -package service -destination mock_test.go github.com/juju/juju/domain/agentprovisioner/service State,Provider
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	model "github.com/juju/juju/core/model"
	envcontext "github.com/juju/juju/environs/envcontext"
	gomock "go.uber.org/mock/gomock"
)

// MockState is a mock of State interface.
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState.
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance.
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// GetModelConfigKeyValues mocks base method.
func (m *MockState) GetModelConfigKeyValues(arg0 context.Context, arg1 ...string) (map[string]string, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetModelConfigKeyValues", varargs...)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModelConfigKeyValues indicates an expected call of GetModelConfigKeyValues.
func (mr *MockStateMockRecorder) GetModelConfigKeyValues(arg0 any, arg1 ...any) *MockStateGetModelConfigKeyValuesCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelConfigKeyValues", reflect.TypeOf((*MockState)(nil).GetModelConfigKeyValues), varargs...)
	return &MockStateGetModelConfigKeyValuesCall{Call: call}
}

// MockStateGetModelConfigKeyValuesCall wrap *gomock.Call
type MockStateGetModelConfigKeyValuesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetModelConfigKeyValuesCall) Return(arg0 map[string]string, arg1 error) *MockStateGetModelConfigKeyValuesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetModelConfigKeyValuesCall) Do(f func(context.Context, ...string) (map[string]string, error)) *MockStateGetModelConfigKeyValuesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetModelConfigKeyValuesCall) DoAndReturn(f func(context.Context, ...string) (map[string]string, error)) *MockStateGetModelConfigKeyValuesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ModelID mocks base method.
func (m *MockState) ModelID(arg0 context.Context) (model.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelID", arg0)
	ret0, _ := ret[0].(model.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelID indicates an expected call of ModelID.
func (mr *MockStateMockRecorder) ModelID(arg0 any) *MockStateModelIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelID", reflect.TypeOf((*MockState)(nil).ModelID), arg0)
	return &MockStateModelIDCall{Call: call}
}

// MockStateModelIDCall wrap *gomock.Call
type MockStateModelIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateModelIDCall) Return(arg0 model.UUID, arg1 error) *MockStateModelIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateModelIDCall) Do(f func(context.Context) (model.UUID, error)) *MockStateModelIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateModelIDCall) DoAndReturn(f func(context.Context) (model.UUID, error)) *MockStateModelIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockProvider is a mock of Provider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider.
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// SupportsContainerAddresses mocks base method.
func (m *MockProvider) SupportsContainerAddresses(arg0 envcontext.ProviderCallContext) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SupportsContainerAddresses", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SupportsContainerAddresses indicates an expected call of SupportsContainerAddresses.
func (mr *MockProviderMockRecorder) SupportsContainerAddresses(arg0 any) *MockProviderSupportsContainerAddressesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupportsContainerAddresses", reflect.TypeOf((*MockProvider)(nil).SupportsContainerAddresses), arg0)
	return &MockProviderSupportsContainerAddressesCall{Call: call}
}

// MockProviderSupportsContainerAddressesCall wrap *gomock.Call
type MockProviderSupportsContainerAddressesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProviderSupportsContainerAddressesCall) Return(arg0 bool, arg1 error) *MockProviderSupportsContainerAddressesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProviderSupportsContainerAddressesCall) Do(f func(envcontext.ProviderCallContext) (bool, error)) *MockProviderSupportsContainerAddressesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProviderSupportsContainerAddressesCall) DoAndReturn(f func(envcontext.ProviderCallContext) (bool, error)) *MockProviderSupportsContainerAddressesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}