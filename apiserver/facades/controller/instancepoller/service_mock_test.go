// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/controller/instancepoller (interfaces: ControllerConfigService,NetworkService,MachineService)
//
// Generated by this command:
//
//	mockgen -typed -package instancepoller_test -destination service_mock_test.go github.com/juju/juju/apiserver/facades/controller/instancepoller ControllerConfigService,NetworkService,MachineService
//

// Package instancepoller_test is a generated GoMock package.
package instancepoller_test

import (
	context "context"
	reflect "reflect"

	controller "github.com/juju/juju/controller"
	machine "github.com/juju/juju/core/machine"
	network "github.com/juju/juju/core/network"
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

// MockNetworkService is a mock of NetworkService interface.
type MockNetworkService struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkServiceMockRecorder
}

// MockNetworkServiceMockRecorder is the mock recorder for MockNetworkService.
type MockNetworkServiceMockRecorder struct {
	mock *MockNetworkService
}

// NewMockNetworkService creates a new mock instance.
func NewMockNetworkService(ctrl *gomock.Controller) *MockNetworkService {
	mock := &MockNetworkService{ctrl: ctrl}
	mock.recorder = &MockNetworkServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkService) EXPECT() *MockNetworkServiceMockRecorder {
	return m.recorder
}

// GetAllSpaces mocks base method.
func (m *MockNetworkService) GetAllSpaces(arg0 context.Context) (network.SpaceInfos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSpaces", arg0)
	ret0, _ := ret[0].(network.SpaceInfos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSpaces indicates an expected call of GetAllSpaces.
func (mr *MockNetworkServiceMockRecorder) GetAllSpaces(arg0 any) *MockNetworkServiceGetAllSpacesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSpaces", reflect.TypeOf((*MockNetworkService)(nil).GetAllSpaces), arg0)
	return &MockNetworkServiceGetAllSpacesCall{Call: call}
}

// MockNetworkServiceGetAllSpacesCall wrap *gomock.Call
type MockNetworkServiceGetAllSpacesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockNetworkServiceGetAllSpacesCall) Return(arg0 network.SpaceInfos, arg1 error) *MockNetworkServiceGetAllSpacesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockNetworkServiceGetAllSpacesCall) Do(f func(context.Context) (network.SpaceInfos, error)) *MockNetworkServiceGetAllSpacesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockNetworkServiceGetAllSpacesCall) DoAndReturn(f func(context.Context) (network.SpaceInfos, error)) *MockNetworkServiceGetAllSpacesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockMachineService is a mock of MachineService interface.
type MockMachineService struct {
	ctrl     *gomock.Controller
	recorder *MockMachineServiceMockRecorder
}

// MockMachineServiceMockRecorder is the mock recorder for MockMachineService.
type MockMachineServiceMockRecorder struct {
	mock *MockMachineService
}

// NewMockMachineService creates a new mock instance.
func NewMockMachineService(ctrl *gomock.Controller) *MockMachineService {
	mock := &MockMachineService{ctrl: ctrl}
	mock.recorder = &MockMachineServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachineService) EXPECT() *MockMachineServiceMockRecorder {
	return m.recorder
}

// EnsureDeadMachine mocks base method.
func (m *MockMachineService) EnsureDeadMachine(arg0 context.Context, arg1 machine.Name) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureDeadMachine", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureDeadMachine indicates an expected call of EnsureDeadMachine.
func (mr *MockMachineServiceMockRecorder) EnsureDeadMachine(arg0, arg1 any) *MockMachineServiceEnsureDeadMachineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureDeadMachine", reflect.TypeOf((*MockMachineService)(nil).EnsureDeadMachine), arg0, arg1)
	return &MockMachineServiceEnsureDeadMachineCall{Call: call}
}

// MockMachineServiceEnsureDeadMachineCall wrap *gomock.Call
type MockMachineServiceEnsureDeadMachineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineServiceEnsureDeadMachineCall) Return(arg0 error) *MockMachineServiceEnsureDeadMachineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineServiceEnsureDeadMachineCall) Do(f func(context.Context, machine.Name) error) *MockMachineServiceEnsureDeadMachineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineServiceEnsureDeadMachineCall) DoAndReturn(f func(context.Context, machine.Name) error) *MockMachineServiceEnsureDeadMachineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetMachineUUID mocks base method.
func (m *MockMachineService) GetMachineUUID(arg0 context.Context, arg1 machine.Name) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachineUUID", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMachineUUID indicates an expected call of GetMachineUUID.
func (mr *MockMachineServiceMockRecorder) GetMachineUUID(arg0, arg1 any) *MockMachineServiceGetMachineUUIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachineUUID", reflect.TypeOf((*MockMachineService)(nil).GetMachineUUID), arg0, arg1)
	return &MockMachineServiceGetMachineUUIDCall{Call: call}
}

// MockMachineServiceGetMachineUUIDCall wrap *gomock.Call
type MockMachineServiceGetMachineUUIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineServiceGetMachineUUIDCall) Return(arg0 string, arg1 error) *MockMachineServiceGetMachineUUIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineServiceGetMachineUUIDCall) Do(f func(context.Context, machine.Name) (string, error)) *MockMachineServiceGetMachineUUIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineServiceGetMachineUUIDCall) DoAndReturn(f func(context.Context, machine.Name) (string, error)) *MockMachineServiceGetMachineUUIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// InstanceID mocks base method.
func (m *MockMachineService) InstanceID(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceID", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceID indicates an expected call of InstanceID.
func (mr *MockMachineServiceMockRecorder) InstanceID(arg0, arg1 any) *MockMachineServiceInstanceIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceID", reflect.TypeOf((*MockMachineService)(nil).InstanceID), arg0, arg1)
	return &MockMachineServiceInstanceIDCall{Call: call}
}

// MockMachineServiceInstanceIDCall wrap *gomock.Call
type MockMachineServiceInstanceIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineServiceInstanceIDCall) Return(arg0 string, arg1 error) *MockMachineServiceInstanceIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineServiceInstanceIDCall) Do(f func(context.Context, string) (string, error)) *MockMachineServiceInstanceIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineServiceInstanceIDCall) DoAndReturn(f func(context.Context, string) (string, error)) *MockMachineServiceInstanceIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// InstanceIDAndName mocks base method.
func (m *MockMachineService) InstanceIDAndName(arg0 context.Context, arg1 string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceIDAndName", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// InstanceIDAndName indicates an expected call of InstanceIDAndName.
func (mr *MockMachineServiceMockRecorder) InstanceIDAndName(arg0, arg1 any) *MockMachineServiceInstanceIDAndNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceIDAndName", reflect.TypeOf((*MockMachineService)(nil).InstanceIDAndName), arg0, arg1)
	return &MockMachineServiceInstanceIDAndNameCall{Call: call}
}

// MockMachineServiceInstanceIDAndNameCall wrap *gomock.Call
type MockMachineServiceInstanceIDAndNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineServiceInstanceIDAndNameCall) Return(arg0, arg1 string, arg2 error) *MockMachineServiceInstanceIDAndNameCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineServiceInstanceIDAndNameCall) Do(f func(context.Context, string) (string, string, error)) *MockMachineServiceInstanceIDAndNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineServiceInstanceIDAndNameCall) DoAndReturn(f func(context.Context, string) (string, string, error)) *MockMachineServiceInstanceIDAndNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
