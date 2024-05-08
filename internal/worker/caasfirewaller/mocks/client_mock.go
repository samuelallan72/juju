// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/worker/caasfirewaller (interfaces: Client,CAASFirewallerAPI,LifeGetter)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/client_mock.go github.com/juju/juju/internal/worker/caasfirewaller Client,CAASFirewallerAPI,LifeGetter
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	charms "github.com/juju/juju/api/common/charms"
	config "github.com/juju/juju/core/config"
	life "github.com/juju/juju/core/life"
	network "github.com/juju/juju/core/network"
	watcher "github.com/juju/juju/core/watcher"
	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ApplicationCharmInfo mocks base method.
func (m *MockClient) ApplicationCharmInfo(arg0 string) (*charms.CharmInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationCharmInfo", arg0)
	ret0, _ := ret[0].(*charms.CharmInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationCharmInfo indicates an expected call of ApplicationCharmInfo.
func (mr *MockClientMockRecorder) ApplicationCharmInfo(arg0 any) *MockClientApplicationCharmInfoCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationCharmInfo", reflect.TypeOf((*MockClient)(nil).ApplicationCharmInfo), arg0)
	return &MockClientApplicationCharmInfoCall{Call: call}
}

// MockClientApplicationCharmInfoCall wrap *gomock.Call
type MockClientApplicationCharmInfoCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientApplicationCharmInfoCall) Return(arg0 *charms.CharmInfo, arg1 error) *MockClientApplicationCharmInfoCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientApplicationCharmInfoCall) Do(f func(string) (*charms.CharmInfo, error)) *MockClientApplicationCharmInfoCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientApplicationCharmInfoCall) DoAndReturn(f func(string) (*charms.CharmInfo, error)) *MockClientApplicationCharmInfoCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ApplicationConfig mocks base method.
func (m *MockClient) ApplicationConfig(arg0 string) (config.ConfigAttributes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationConfig", arg0)
	ret0, _ := ret[0].(config.ConfigAttributes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationConfig indicates an expected call of ApplicationConfig.
func (mr *MockClientMockRecorder) ApplicationConfig(arg0 any) *MockClientApplicationConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationConfig", reflect.TypeOf((*MockClient)(nil).ApplicationConfig), arg0)
	return &MockClientApplicationConfigCall{Call: call}
}

// MockClientApplicationConfigCall wrap *gomock.Call
type MockClientApplicationConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientApplicationConfigCall) Return(arg0 config.ConfigAttributes, arg1 error) *MockClientApplicationConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientApplicationConfigCall) Do(f func(string) (config.ConfigAttributes, error)) *MockClientApplicationConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientApplicationConfigCall) DoAndReturn(f func(string) (config.ConfigAttributes, error)) *MockClientApplicationConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetOpenedPorts mocks base method.
func (m *MockClient) GetOpenedPorts(arg0 string) (network.GroupedPortRanges, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOpenedPorts", arg0)
	ret0, _ := ret[0].(network.GroupedPortRanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOpenedPorts indicates an expected call of GetOpenedPorts.
func (mr *MockClientMockRecorder) GetOpenedPorts(arg0 any) *MockClientGetOpenedPortsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOpenedPorts", reflect.TypeOf((*MockClient)(nil).GetOpenedPorts), arg0)
	return &MockClientGetOpenedPortsCall{Call: call}
}

// MockClientGetOpenedPortsCall wrap *gomock.Call
type MockClientGetOpenedPortsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientGetOpenedPortsCall) Return(arg0 network.GroupedPortRanges, arg1 error) *MockClientGetOpenedPortsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientGetOpenedPortsCall) Do(f func(string) (network.GroupedPortRanges, error)) *MockClientGetOpenedPortsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientGetOpenedPortsCall) DoAndReturn(f func(string) (network.GroupedPortRanges, error)) *MockClientGetOpenedPortsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsExposed mocks base method.
func (m *MockClient) IsExposed(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExposed", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsExposed indicates an expected call of IsExposed.
func (mr *MockClientMockRecorder) IsExposed(arg0 any) *MockClientIsExposedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExposed", reflect.TypeOf((*MockClient)(nil).IsExposed), arg0)
	return &MockClientIsExposedCall{Call: call}
}

// MockClientIsExposedCall wrap *gomock.Call
type MockClientIsExposedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientIsExposedCall) Return(arg0 bool, arg1 error) *MockClientIsExposedCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientIsExposedCall) Do(f func(string) (bool, error)) *MockClientIsExposedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientIsExposedCall) DoAndReturn(f func(string) (bool, error)) *MockClientIsExposedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Life mocks base method.
func (m *MockClient) Life(arg0 string) (life.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Life", arg0)
	ret0, _ := ret[0].(life.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Life indicates an expected call of Life.
func (mr *MockClientMockRecorder) Life(arg0 any) *MockClientLifeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Life", reflect.TypeOf((*MockClient)(nil).Life), arg0)
	return &MockClientLifeCall{Call: call}
}

// MockClientLifeCall wrap *gomock.Call
type MockClientLifeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientLifeCall) Return(arg0 life.Value, arg1 error) *MockClientLifeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientLifeCall) Do(f func(string) (life.Value, error)) *MockClientLifeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientLifeCall) DoAndReturn(f func(string) (life.Value, error)) *MockClientLifeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchApplication mocks base method.
func (m *MockClient) WatchApplication(arg0 string) (watcher.Watcher[struct{}], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchApplication", arg0)
	ret0, _ := ret[0].(watcher.Watcher[struct{}])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchApplication indicates an expected call of WatchApplication.
func (mr *MockClientMockRecorder) WatchApplication(arg0 any) *MockClientWatchApplicationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchApplication", reflect.TypeOf((*MockClient)(nil).WatchApplication), arg0)
	return &MockClientWatchApplicationCall{Call: call}
}

// MockClientWatchApplicationCall wrap *gomock.Call
type MockClientWatchApplicationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientWatchApplicationCall) Return(arg0 watcher.Watcher[struct{}], arg1 error) *MockClientWatchApplicationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientWatchApplicationCall) Do(f func(string) (watcher.Watcher[struct{}], error)) *MockClientWatchApplicationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientWatchApplicationCall) DoAndReturn(f func(string) (watcher.Watcher[struct{}], error)) *MockClientWatchApplicationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchApplications mocks base method.
func (m *MockClient) WatchApplications() (watcher.Watcher[[]string], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchApplications")
	ret0, _ := ret[0].(watcher.Watcher[[]string])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchApplications indicates an expected call of WatchApplications.
func (mr *MockClientMockRecorder) WatchApplications() *MockClientWatchApplicationsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchApplications", reflect.TypeOf((*MockClient)(nil).WatchApplications))
	return &MockClientWatchApplicationsCall{Call: call}
}

// MockClientWatchApplicationsCall wrap *gomock.Call
type MockClientWatchApplicationsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientWatchApplicationsCall) Return(arg0 watcher.Watcher[[]string], arg1 error) *MockClientWatchApplicationsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientWatchApplicationsCall) Do(f func() (watcher.Watcher[[]string], error)) *MockClientWatchApplicationsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientWatchApplicationsCall) DoAndReturn(f func() (watcher.Watcher[[]string], error)) *MockClientWatchApplicationsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchOpenedPorts mocks base method.
func (m *MockClient) WatchOpenedPorts() (watcher.Watcher[[]string], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchOpenedPorts")
	ret0, _ := ret[0].(watcher.Watcher[[]string])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchOpenedPorts indicates an expected call of WatchOpenedPorts.
func (mr *MockClientMockRecorder) WatchOpenedPorts() *MockClientWatchOpenedPortsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchOpenedPorts", reflect.TypeOf((*MockClient)(nil).WatchOpenedPorts))
	return &MockClientWatchOpenedPortsCall{Call: call}
}

// MockClientWatchOpenedPortsCall wrap *gomock.Call
type MockClientWatchOpenedPortsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientWatchOpenedPortsCall) Return(arg0 watcher.Watcher[[]string], arg1 error) *MockClientWatchOpenedPortsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientWatchOpenedPortsCall) Do(f func() (watcher.Watcher[[]string], error)) *MockClientWatchOpenedPortsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientWatchOpenedPortsCall) DoAndReturn(f func() (watcher.Watcher[[]string], error)) *MockClientWatchOpenedPortsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockCAASFirewallerAPI is a mock of CAASFirewallerAPI interface.
type MockCAASFirewallerAPI struct {
	ctrl     *gomock.Controller
	recorder *MockCAASFirewallerAPIMockRecorder
}

// MockCAASFirewallerAPIMockRecorder is the mock recorder for MockCAASFirewallerAPI.
type MockCAASFirewallerAPIMockRecorder struct {
	mock *MockCAASFirewallerAPI
}

// NewMockCAASFirewallerAPI creates a new mock instance.
func NewMockCAASFirewallerAPI(ctrl *gomock.Controller) *MockCAASFirewallerAPI {
	mock := &MockCAASFirewallerAPI{ctrl: ctrl}
	mock.recorder = &MockCAASFirewallerAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCAASFirewallerAPI) EXPECT() *MockCAASFirewallerAPIMockRecorder {
	return m.recorder
}

// ApplicationCharmInfo mocks base method.
func (m *MockCAASFirewallerAPI) ApplicationCharmInfo(arg0 string) (*charms.CharmInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationCharmInfo", arg0)
	ret0, _ := ret[0].(*charms.CharmInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationCharmInfo indicates an expected call of ApplicationCharmInfo.
func (mr *MockCAASFirewallerAPIMockRecorder) ApplicationCharmInfo(arg0 any) *MockCAASFirewallerAPIApplicationCharmInfoCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationCharmInfo", reflect.TypeOf((*MockCAASFirewallerAPI)(nil).ApplicationCharmInfo), arg0)
	return &MockCAASFirewallerAPIApplicationCharmInfoCall{Call: call}
}

// MockCAASFirewallerAPIApplicationCharmInfoCall wrap *gomock.Call
type MockCAASFirewallerAPIApplicationCharmInfoCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCAASFirewallerAPIApplicationCharmInfoCall) Return(arg0 *charms.CharmInfo, arg1 error) *MockCAASFirewallerAPIApplicationCharmInfoCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCAASFirewallerAPIApplicationCharmInfoCall) Do(f func(string) (*charms.CharmInfo, error)) *MockCAASFirewallerAPIApplicationCharmInfoCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCAASFirewallerAPIApplicationCharmInfoCall) DoAndReturn(f func(string) (*charms.CharmInfo, error)) *MockCAASFirewallerAPIApplicationCharmInfoCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ApplicationConfig mocks base method.
func (m *MockCAASFirewallerAPI) ApplicationConfig(arg0 string) (config.ConfigAttributes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationConfig", arg0)
	ret0, _ := ret[0].(config.ConfigAttributes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationConfig indicates an expected call of ApplicationConfig.
func (mr *MockCAASFirewallerAPIMockRecorder) ApplicationConfig(arg0 any) *MockCAASFirewallerAPIApplicationConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationConfig", reflect.TypeOf((*MockCAASFirewallerAPI)(nil).ApplicationConfig), arg0)
	return &MockCAASFirewallerAPIApplicationConfigCall{Call: call}
}

// MockCAASFirewallerAPIApplicationConfigCall wrap *gomock.Call
type MockCAASFirewallerAPIApplicationConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCAASFirewallerAPIApplicationConfigCall) Return(arg0 config.ConfigAttributes, arg1 error) *MockCAASFirewallerAPIApplicationConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCAASFirewallerAPIApplicationConfigCall) Do(f func(string) (config.ConfigAttributes, error)) *MockCAASFirewallerAPIApplicationConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCAASFirewallerAPIApplicationConfigCall) DoAndReturn(f func(string) (config.ConfigAttributes, error)) *MockCAASFirewallerAPIApplicationConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetOpenedPorts mocks base method.
func (m *MockCAASFirewallerAPI) GetOpenedPorts(arg0 string) (network.GroupedPortRanges, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOpenedPorts", arg0)
	ret0, _ := ret[0].(network.GroupedPortRanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOpenedPorts indicates an expected call of GetOpenedPorts.
func (mr *MockCAASFirewallerAPIMockRecorder) GetOpenedPorts(arg0 any) *MockCAASFirewallerAPIGetOpenedPortsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOpenedPorts", reflect.TypeOf((*MockCAASFirewallerAPI)(nil).GetOpenedPorts), arg0)
	return &MockCAASFirewallerAPIGetOpenedPortsCall{Call: call}
}

// MockCAASFirewallerAPIGetOpenedPortsCall wrap *gomock.Call
type MockCAASFirewallerAPIGetOpenedPortsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCAASFirewallerAPIGetOpenedPortsCall) Return(arg0 network.GroupedPortRanges, arg1 error) *MockCAASFirewallerAPIGetOpenedPortsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCAASFirewallerAPIGetOpenedPortsCall) Do(f func(string) (network.GroupedPortRanges, error)) *MockCAASFirewallerAPIGetOpenedPortsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCAASFirewallerAPIGetOpenedPortsCall) DoAndReturn(f func(string) (network.GroupedPortRanges, error)) *MockCAASFirewallerAPIGetOpenedPortsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsExposed mocks base method.
func (m *MockCAASFirewallerAPI) IsExposed(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExposed", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsExposed indicates an expected call of IsExposed.
func (mr *MockCAASFirewallerAPIMockRecorder) IsExposed(arg0 any) *MockCAASFirewallerAPIIsExposedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExposed", reflect.TypeOf((*MockCAASFirewallerAPI)(nil).IsExposed), arg0)
	return &MockCAASFirewallerAPIIsExposedCall{Call: call}
}

// MockCAASFirewallerAPIIsExposedCall wrap *gomock.Call
type MockCAASFirewallerAPIIsExposedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCAASFirewallerAPIIsExposedCall) Return(arg0 bool, arg1 error) *MockCAASFirewallerAPIIsExposedCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCAASFirewallerAPIIsExposedCall) Do(f func(string) (bool, error)) *MockCAASFirewallerAPIIsExposedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCAASFirewallerAPIIsExposedCall) DoAndReturn(f func(string) (bool, error)) *MockCAASFirewallerAPIIsExposedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchApplication mocks base method.
func (m *MockCAASFirewallerAPI) WatchApplication(arg0 string) (watcher.Watcher[struct{}], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchApplication", arg0)
	ret0, _ := ret[0].(watcher.Watcher[struct{}])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchApplication indicates an expected call of WatchApplication.
func (mr *MockCAASFirewallerAPIMockRecorder) WatchApplication(arg0 any) *MockCAASFirewallerAPIWatchApplicationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchApplication", reflect.TypeOf((*MockCAASFirewallerAPI)(nil).WatchApplication), arg0)
	return &MockCAASFirewallerAPIWatchApplicationCall{Call: call}
}

// MockCAASFirewallerAPIWatchApplicationCall wrap *gomock.Call
type MockCAASFirewallerAPIWatchApplicationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCAASFirewallerAPIWatchApplicationCall) Return(arg0 watcher.Watcher[struct{}], arg1 error) *MockCAASFirewallerAPIWatchApplicationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCAASFirewallerAPIWatchApplicationCall) Do(f func(string) (watcher.Watcher[struct{}], error)) *MockCAASFirewallerAPIWatchApplicationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCAASFirewallerAPIWatchApplicationCall) DoAndReturn(f func(string) (watcher.Watcher[struct{}], error)) *MockCAASFirewallerAPIWatchApplicationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchApplications mocks base method.
func (m *MockCAASFirewallerAPI) WatchApplications() (watcher.Watcher[[]string], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchApplications")
	ret0, _ := ret[0].(watcher.Watcher[[]string])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchApplications indicates an expected call of WatchApplications.
func (mr *MockCAASFirewallerAPIMockRecorder) WatchApplications() *MockCAASFirewallerAPIWatchApplicationsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchApplications", reflect.TypeOf((*MockCAASFirewallerAPI)(nil).WatchApplications))
	return &MockCAASFirewallerAPIWatchApplicationsCall{Call: call}
}

// MockCAASFirewallerAPIWatchApplicationsCall wrap *gomock.Call
type MockCAASFirewallerAPIWatchApplicationsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCAASFirewallerAPIWatchApplicationsCall) Return(arg0 watcher.Watcher[[]string], arg1 error) *MockCAASFirewallerAPIWatchApplicationsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCAASFirewallerAPIWatchApplicationsCall) Do(f func() (watcher.Watcher[[]string], error)) *MockCAASFirewallerAPIWatchApplicationsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCAASFirewallerAPIWatchApplicationsCall) DoAndReturn(f func() (watcher.Watcher[[]string], error)) *MockCAASFirewallerAPIWatchApplicationsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchOpenedPorts mocks base method.
func (m *MockCAASFirewallerAPI) WatchOpenedPorts() (watcher.Watcher[[]string], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchOpenedPorts")
	ret0, _ := ret[0].(watcher.Watcher[[]string])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchOpenedPorts indicates an expected call of WatchOpenedPorts.
func (mr *MockCAASFirewallerAPIMockRecorder) WatchOpenedPorts() *MockCAASFirewallerAPIWatchOpenedPortsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchOpenedPorts", reflect.TypeOf((*MockCAASFirewallerAPI)(nil).WatchOpenedPorts))
	return &MockCAASFirewallerAPIWatchOpenedPortsCall{Call: call}
}

// MockCAASFirewallerAPIWatchOpenedPortsCall wrap *gomock.Call
type MockCAASFirewallerAPIWatchOpenedPortsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCAASFirewallerAPIWatchOpenedPortsCall) Return(arg0 watcher.Watcher[[]string], arg1 error) *MockCAASFirewallerAPIWatchOpenedPortsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCAASFirewallerAPIWatchOpenedPortsCall) Do(f func() (watcher.Watcher[[]string], error)) *MockCAASFirewallerAPIWatchOpenedPortsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCAASFirewallerAPIWatchOpenedPortsCall) DoAndReturn(f func() (watcher.Watcher[[]string], error)) *MockCAASFirewallerAPIWatchOpenedPortsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockLifeGetter is a mock of LifeGetter interface.
type MockLifeGetter struct {
	ctrl     *gomock.Controller
	recorder *MockLifeGetterMockRecorder
}

// MockLifeGetterMockRecorder is the mock recorder for MockLifeGetter.
type MockLifeGetterMockRecorder struct {
	mock *MockLifeGetter
}

// NewMockLifeGetter creates a new mock instance.
func NewMockLifeGetter(ctrl *gomock.Controller) *MockLifeGetter {
	mock := &MockLifeGetter{ctrl: ctrl}
	mock.recorder = &MockLifeGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLifeGetter) EXPECT() *MockLifeGetterMockRecorder {
	return m.recorder
}

// Life mocks base method.
func (m *MockLifeGetter) Life(arg0 string) (life.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Life", arg0)
	ret0, _ := ret[0].(life.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Life indicates an expected call of Life.
func (mr *MockLifeGetterMockRecorder) Life(arg0 any) *MockLifeGetterLifeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Life", reflect.TypeOf((*MockLifeGetter)(nil).Life), arg0)
	return &MockLifeGetterLifeCall{Call: call}
}

// MockLifeGetterLifeCall wrap *gomock.Call
type MockLifeGetterLifeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLifeGetterLifeCall) Return(arg0 life.Value, arg1 error) *MockLifeGetterLifeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLifeGetterLifeCall) Do(f func(string) (life.Value, error)) *MockLifeGetterLifeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLifeGetterLifeCall) DoAndReturn(f func(string) (life.Value, error)) *MockLifeGetterLifeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
